package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/payment"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
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

	// transaction
	transactionStore := dbstore.NewTransactionStore()
	transaction := &store.Transaction{
		BuildTime: 0,                       //buildTime := r.FormValue("time")
		Quantity:  r.FormValue("quantity"), // 	quantity := r.FormValue("quantity")
		Price:     0,                       // price := r.FormValue("price")
		Phone:     r.FormValue("phone"),
	}
	err = transactionStore.CreateTransaction(transaction)
	if err != nil {
		fmt.Println("Error saving saving transaction to db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	// payment
	phone := r.FormValue("phone")
	phone = strings.TrimPrefix(phone, "0")
	phone = "254" + phone
	intPhone, err := strconv.ParseInt(phone, 10, 64)

	paymentProcessor := payment.NewPaymentProcessor(int(intPhone))
	paymentProcessor.InitiateStkPush()

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	return Render(w, r, components.PaymentForm())
}
