package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/payment"
)

func DarajaCallbackHandler(w http.ResponseWriter, r *http.Request) error {
	var callbackResponse payment.StkCallbackResponse
	fmt.Println("IN DARAJA CALLBACK")
	err := json.NewDecoder(r.Body).Decode(&callbackResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}

	stkCallback := callbackResponse.Body.StkCallback

	fmt.Printf("Callback received:\n")
	fmt.Printf("MerchantRequestID: %s\n", stkCallback.MerchantRequestID)
	fmt.Printf("CheckoutRequestID: %s\n", stkCallback.CheckoutRequestID)
	fmt.Printf("ResultCode: %d\n", stkCallback.ResultCode)
	fmt.Printf("ResultDesc: %s\n", stkCallback.ResultDesc)

	// amount mpesa receiptno transaction date phone number
	for _, item := range stkCallback.CallbackMetadata.Item {
		fmt.Printf("%s: %v\n", item.Name, item.Value)
	}

	// get user with phone number
	// select * from orders
	// join users on users.id = orders.user_id
	// where orders.payment_complete=false and
	// orders.phone=user.phone

	// defer res.Body.Close()
	// body, err := io.ReadAll(res.Body)
	// fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	return nil
}
