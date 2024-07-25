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

func DarajaPaymentStatusCallback(w http.ResponseWriter, r *http.Request) error {
	var callbackResponse payment.PaymentStatusCallbackResponse

	err := json.NewDecoder(r.Body).Decode(&callbackResponse)
	if err != nil {
		return err
	}

	callbackResult := callbackResponse.Result

	fmt.Printf("Callback received:\n")
	fmt.Println(callbackResult.TransactionID)
	fmt.Println(callbackResult.ConversationID)
	fmt.Println(callbackResult.OriginatorConversationID)
	fmt.Println(callbackResult.ResultCode)
	fmt.Println(callbackResult.ResultDesc)

	transactionStore := dbstore.NewTransactionStore()
	if callbackResult.ResultCode != 0 {
		transactionStore.UpdateTransactionState(
			callbackResult.TransactionID,
			fmt.Sprint(store.PaymentFailed),
		)

		return errors.New(fmt.Sprintf(
			"Transaction ID: %s  Error: %s",
			callbackResult.TransactionID,
			callbackResult.ResultDesc,
		))
	}

	err = transactionStore.UpdateTransactionState(
		callbackResult.TransactionID,
		fmt.Sprint(store.PaymentComplete),
	)

	w.WriteHeader(http.StatusOK)
	return err
}
