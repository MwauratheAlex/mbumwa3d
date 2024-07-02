package handlers

import (
	"fmt"
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"net/http"
)

func PostPrint(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if ok == false {
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, components.UnauthorizedFormEror())
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
	msg, err := filestore.SaveToDisk(file, handler.Filename)

	if err != nil {
		fmt.Println(msg, err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

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
		BuildTime:       0, // to be calculated
		Quantity:        r.FormValue("quantity"),
		Price:           0, // to be calculated
		Phone:           r.FormValue("phone"),
		PaymentComplete: false,
		Status:          "Completed",
	}
	err = orderStore.CreateOrder(order)
	if err != nil {
		fmt.Println("Error saving saving order to db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	// add order to cart

	// fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	return Render(w, r, components.PaymentForm(
		fmt.Sprintf("%.2f", order.Price),
		fmt.Sprintf("%d", order.BuildTime),
	))
}
