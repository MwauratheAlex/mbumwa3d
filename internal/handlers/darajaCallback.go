package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/payment"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

func DarajaCallbackHandler(w http.ResponseWriter, r *http.Request) error {
	var callbackResponse payment.StkCallbackResponse

	err := json.NewDecoder(r.Body).Decode(&callbackResponse)
	if err != nil {
		return err
	}

	stkCallback := callbackResponse.Body.StkCallback

	fmt.Printf("Callback received:\n")

	transactionStore := dbstore.NewTransactionStore()
	if stkCallback.ResultCode != 0 {
		transactionStore.UpdateTransactionState(
			stkCallback.CheckoutRequestID,
			fmt.Sprint(store.PaymentFailed),
		)

		return errors.New(fmt.Sprintf(
			"Transaction ID: %s  Error: %s",
			stkCallback.CheckoutRequestID,
			stkCallback.ResultDesc,
		))
	}

	err = transactionStore.UpdateTransactionState(
		stkCallback.CheckoutRequestID,
		fmt.Sprint(store.PaymentComplete),
	)

	w.WriteHeader(http.StatusOK)
	return err
}
