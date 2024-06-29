package handlers

import (
	"fmt"
	"net/http"
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
		fmt.Println("Unable to pass form data")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return Render(w, r, components.UploadFormError())
	}

	fmt.Println("file uploaded")
	// do file stuff

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

	w.WriteHeader(http.StatusOK)
	return Render(w, r, components.PaymentForm())
}
