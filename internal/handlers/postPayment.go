package handlers

import (
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"net/http"
)

func PostPayment(w http.ResponseWriter, r *http.Request) error {
	_, ok := r.Context().Value(middleware.UserKey).(*store.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, components.UnauthorizedFormEror())
	}
	// get cart with orders
	// add mobile number to card

	// payment
	// phone := r.FormValue("phone")
	// phone = strings.TrimPrefix(phone, "0")
	// phone = "254" + phone
	// intPhone, err := strconv.ParseInt(phone, 10, 64)
	// if err != nil {
	// 	fmt.Println("Error Parsing Phone number", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return Render(w, r, components.UploadFormError("Internal server error."))
	// }

	// paymentProcessor := payment.NewPaymentProcessor(int(intPhone))
	// paymentProcessor.InitiateStkPush()

	// get the cart
	// calculate price

	return nil
}
