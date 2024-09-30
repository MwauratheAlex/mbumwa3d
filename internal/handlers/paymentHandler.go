package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mwaurathealex/mbumwa3d/internal/payment"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

type PaymentHandler struct {
	SessionName      string
	PaymentProcessor *payment.PaymentProcessor
	OrderStore       *dbstore.OrderStore
}

type PaymentHandlerParams struct {
	PaymentProcessor *payment.PaymentProcessor
	OrderStore       *dbstore.OrderStore
	SessionName      string
}

func NewPaymentHandler(params PaymentHandlerParams) *PaymentHandler {
	return &PaymentHandler{
		PaymentProcessor: params.PaymentProcessor,
		SessionName:      params.SessionName,
		OrderStore:       params.OrderStore,
	}
}

func (h *PaymentHandler) Post(w http.ResponseWriter, r *http.Request) error {
	// 7. Show toast notifications

	phone := r.FormValue("phone")
	if phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	printConfig, err := GetSessionPrintConfig(r, h.SessionName)
	err = ValidatePrintConfig(&printConfig)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	printConfig.UserID = uint(userID)

	fmt.Println("HERE", phone)
	w.WriteHeader(http.StatusInternalServerError)

	h.calculateConfigPrice(&printConfig)
	err = h.OrderStore.CreatePrintConfig(&printConfig)
	if err != nil {
		panic(err)
	}

	price := printConfig.Price

	order := &store.Order{
		UserID:          uint(userID),
		PrintConfigID:   printConfig.ID,
		Price:           price,
		PaymentComplete: false,
		Status:          fmt.Sprint(store.AwaitingPayment),
	}
	fmt.Println("h", printConfig.Price, err, userID, order.PrintConfigID)

	err = h.OrderStore.CreateOrder(order)

	intPhone, err := strconv.ParseInt("254"+phone, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(intPhone, "price", int(price))

	price = 1

	_, err = h.PaymentProcessor.InitiateStkPush(
		int(price), int(intPhone),
	)

	if err != nil {
		panic(err)
	}

	// order.CheckoutRequestId = transactionResponse.CheckoutRequestID
	// fmt.Println(transactionResponse)

	// h.OrderStore.Save(order)

	// _, err = h.PaymentProcessor.GetTransactionStatus(
	// 	transactionResponse.CheckoutRequestID,
	// )
	return err
}

func (h *PaymentHandler) calculateConfigPrice(config *store.PrintConfig) {
	basePricePerCubicMM := 0.05
	technologyMultiplier := 1.0
	switch config.Technology {
	case "FDM":
		technologyMultiplier = 1.0
	case "SLA":
		technologyMultiplier = 1.5
	}

	materialMultiplier := 1.0
	switch config.Material {
	case "PLA":
		materialMultiplier = 1.0
	case "ABS":
		materialMultiplier = 1.2
	}

	colorMultiplier := 1.0

	price := config.FileVolume * basePricePerCubicMM * technologyMultiplier *
		materialMultiplier * colorMultiplier * float64(config.Quantity)

	config.Price = price
}

func (h *PaymentHandler) DarajaCallback(w http.ResponseWriter, r *http.Request) error {
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
			fmt.Sprint(""),
		)

		return errors.New(fmt.Sprintf(
			"Transaction ID: %s  Error: %s",
			stkCallback.CheckoutRequestID,
			stkCallback.ResultDesc,
		))
	}

	err = transactionStore.UpdateTransactionState(
		stkCallback.CheckoutRequestID,
		fmt.Sprint(""),
	)

	w.WriteHeader(http.StatusOK)
	return err
}

type PaymentNotificationRes struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             int    `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (h *PaymentHandler) PaymentNotificationCallback(
	w http.ResponseWriter, r *http.Request) {
	var res PaymentNotificationRes
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		panic(err)
	}

	fmt.Println("In payment notify callback")
	fmt.Println("ResponseCode: ", res.OriginatorConversationID)
	fmt.Println("ResponseCode: ", res.ResponseCode)
	fmt.Println("ResponseCode: ", res.ResponseDescription)
}

func (h *PaymentHandler) DarajaPaymentStatusCallback(w http.ResponseWriter, r *http.Request) error {
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
			fmt.Sprint(""),
		)

		return errors.New(fmt.Sprintf(
			"Transaction ID: %s  Error: %s",
			callbackResult.TransactionID,
			callbackResult.ResultDesc,
		))
	}

	err = transactionStore.UpdateTransactionState(
		callbackResult.TransactionID,
		fmt.Sprint(""),
	)

	w.WriteHeader(http.StatusOK)
	return err
}
