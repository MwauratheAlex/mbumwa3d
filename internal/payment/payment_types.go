package payment

import "net/http"

type StkPushRequest struct {
	BusinessShortCode int64  `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int    `json:"Amount"`
	PartyA            int    `json:"PartyA"`
	PartyB            int64  `json:"PartyB"`
	PhoneNumber       int    `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type PaymentProcessor struct {
	DarajaAuthUrl              string
	DarajaStkPushUrl           string
	DarajaTransactionStatusUrl string
	CallBackURL                string
	Client                     *http.Client
	PhoneNumber                int
	BusinessShortCode          int64
	PassKey                    string
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

type TransactionStatusRequest struct {
	CommandID                string
	PartyA                   int64
	IdentifierType           string
	Remarks                  string
	Initiator                string
	SecurityCredential       string
	QueueTimeOutURL          string
	TransactionID            string
	ResultURL                string
	Occasion                 string
	OriginatorConversationID string
}

type TransactionStatusResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

type ResultParameter struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type ResultParameters struct {
	ResultParameter []ResultParameter `json:"ResultParameter"`
}

type ReferenceItem struct {
	Key string `json:"Key"`
}

type ReferenceData struct {
	ReferenceItem ReferenceItem `json:"ReferenceItem"`
}

type Result struct {
	ConversationID           string           `json:"ConversationID"`
	OriginatorConversationID string           `json:"OriginatorConversationID"`
	ReferenceData            ReferenceData    `json:"ReferenceData"`
	ResultCode               int              `json:"ResultCode"`
	ResultDesc               string           `json:"ResultDesc"`
	ResultParameters         ResultParameters `json:"ResultParameters"`
	ResultType               int              `json:"ResultType"`
	TransactionID            string           `json:"TransactionID"`
}

type PaymentStatusCallbackResponse struct {
	Result Result `json:"Result"`
}
