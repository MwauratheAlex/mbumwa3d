package payment

import "net/http"

type StkPushRequest struct {
	BusinessShortCode int    `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int    `json:"Amount"`
	PartyA            int    `json:"PartyA"`
	PartyB            int    `json:"PartyB"`
	PhoneNumber       int    `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type PaymentProcessor struct {
	DarajaAuthUrl    string
	DarajaStkPushUrl string
	Client           *http.Client
	PhoneNumber      int
	CallBackURL      string
}

type TransactionResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

type StkCallbackMetadata struct {
	Item []struct {
		Name  string      `json:"Name"`
		Value interface{} `json:"Value,omitempty"`
	} `json:"Item"`
}

type StkCallback struct {
	MerchantRequestID string              `json:"MerchantRequestID"`
	CheckoutRequestID string              `json:"CheckoutRequestID"`
	ResultCode        int                 `json:"ResultCode"`
	ResultDesc        string              `json:"ResultDesc"`
	CallbackMetadata  StkCallbackMetadata `json:"CallbackMetadata"`
}

type StkCallbackBody struct {
	StkCallback StkCallback `json:"stkCallback"`
}

type StkCallbackResponse struct {
	Body StkCallbackBody `json:"Body"`
}
