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

func PostPayment(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, components.UnauthorizedFormEror())
	}

	phone := r.FormValue("phone")

	cartStore := dbstore.NewCartStore(user.ID)
	cart := cartStore.GetCartByUserId()

	transaction := cart.Transaction
	transaction.Phone = phone

	// payment
	phone = strings.TrimPrefix(phone, "0")
	phone = "254" + phone
	intPhone, err := strconv.ParseInt(phone, 10, 64)
	if err != nil {
		fmt.Println("Error Parsing Phone number", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Internal server error."))
	}

	paymentProcessor := payment.NewPaymentProcessor(int(intPhone))
	// calculate sum of all prices
	transactionResponse, err := paymentProcessor.InitiateStkPush(1)
	if err != nil {
		fmt.Println("Transaction failed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, components.UploadFormError("Transaction failed."))
	}

	// save the transaction and wait for callback
	transaction.CheckoutRequestId = transactionResponse.CheckoutRequestID
	transaction.PaymentStatus = fmt.Sprint(store.ProcessingPayment)

	transactionStore := dbstore.NewTransactionStore(user.ID)
	transactionStore.SaveTransaction(transaction)
	cartStore.ClearCart(cart)

	return Render(w, r, components.UploadForm())
}
