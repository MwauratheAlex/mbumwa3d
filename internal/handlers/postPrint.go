package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/stl"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
)

func PostPrint(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if ok == false {
		w.Header().Add("HX-Trigger", GetToastPayload("PrintError", "unauthorized"))
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		fmt.Println("Unable to parse form data:", err)
		return Render(w, r, components.UploadFormError("Invalid form data."))
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error getting file", err)
		w.WriteHeader(http.StatusBadRequest)
		return Render(w, r, components.UploadFormError(
			"Please upload a file before submitting."),
		)
	}

	// file
	filestore := dbstore.NewFileStore()
	dstPath, msg, err := filestore.SaveToDisk(file, handler.Filename)

	if err != nil {
		fmt.Println(msg, err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	// price
	stlCalc, err := stl.NewSTLCalc(dstPath)
	if err != nil {
		fmt.Println("Error creating stl calc: ", err)
	}
	defer stlCalc.Close()
	price, err := stlCalc.CalculatePrice("FDM", "ABS", 1)
	if err != nil {
		fmt.Println("Error calculating price: ", err)
	}

	// file to db
	dbfile := &store.File{
		UserID:     user.ID,
		LocalPath:  handler.Filename,
		FileName:   handler.Filename,
		Technology: r.FormValue("technology"),
		Color:      r.FormValue("color"),
	}
	err = filestore.SaveFileToDB(dbfile)
	if err != nil {
		fmt.Println("Error saving file to db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	// order
	orderStore := dbstore.NewOrderStore()
	order := &store.Order{
		UserID:          user.ID,
		FileID:          dbfile.ID,
		BuildTime:       72, // to be calculated
		Quantity:        r.FormValue("quantity"),
		Price:           price,
		PaymentComplete: false,
		Status:          fmt.Sprint(store.Reviewing),
		PrintStatus:     fmt.Sprint(store.Available),
	}
	err = orderStore.CreateOrder(order)
	if err != nil {
		fmt.Println("Error saving saving order to db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}
	fmt.Println("Status", order.Status)

	// add order to cart
	cartStore := dbstore.NewCartStore(user.ID)
	cart := cartStore.GetCartByUserId()

	if cart.Transaction == nil {
		cart.Transaction = &store.Transaction{
			UserID:        user.ID,
			PaymentStatus: fmt.Sprint(store.AwaitingPayment),
		}
	}
	cart.Transaction.Orders = append(cart.Transaction.Orders, *order)
	cartStore.SaveCart(cart)
	cart.TransactionID = cart.Transaction.ID

	cartStore.SaveCart(cart)
	err = cartStore.SaveCart(cart)

	if err != nil {
		fmt.Println("Error adding to cart", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	return Render(w, r, components.PaymentForm(
		fmt.Sprintf("%.2f", order.Price),
		fmt.Sprintf("%d", order.BuildTime),
	))
}
