package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/payment"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
)

func PostPrint(w http.ResponseWriter, r *http.Request) error {
	_, ok := r.Context().Value(middleware.UserKey).(*store.User)

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

	defer file.Close()
	// write file to disk
	dir := "./model_storage"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}
	dstPath := filepath.Join(dir, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		fmt.Println("Error creating file on disk:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	// store data in db
	technology := r.FormValue("technology")
	color := r.FormValue("color")
	buildTime := r.FormValue("time")
	quantity := r.FormValue("quantity")
	price := r.FormValue("price")
	phone := r.FormValue("phone")
	phone = strings.TrimPrefix(phone, "0")
	phone = "254" + phone
	intPhone, err := strconv.ParseInt(phone, 10, 64)

	paymentProcessor := payment.NewPaymentProcessor(int(intPhone))
	paymentProcessor.InitiateStkPush()

	fmt.Println("tech: ", technology, "Color: ", color, "time: ",
		buildTime, "qty: ", quantity, "price: ", price, "phone: ", phone)
	fmt.Println(file)

	return Render(w, r, components.PaymentForm())
}
