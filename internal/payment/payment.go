package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func NewPaymentProcessor(phoneNumber int) *PaymentProcessor {
	return &PaymentProcessor{
		DarajaAuthUrl:    os.Getenv("DARAJA_AUTH_URL"),
		DarajaStkPushUrl: os.Getenv("DARAJA_STK_PUSH_URL"),
		Client:           &http.Client{},
		PhoneNumber:      phoneNumber,
	}
}

func (p *PaymentProcessor) GetAuthToken() (string, error) {
	req, err := http.NewRequest("GET", p.DarajaAuthUrl, nil)
	if err != nil {
		return "", err
	}
	darajaAuthHeader := os.Getenv("DARAJA_AUTH_HEADER")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", darajaAuthHeader)

	res, err := p.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func (p *PaymentProcessor) GeneratePassword(shortcode int, passkey string) string {
	timestamp := time.Now().Format("20060102150405")
	data := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
	password := base64.StdEncoding.EncodeToString([]byte(data))
	return password
}

func (p *PaymentProcessor) InitiateStkPush(amount int) {
	businessShortCode64, err := strconv.ParseInt(os.Getenv("DARAJA_SHORTCODE"), 10, 64)
	if err != nil {
		println(err)
	}
	passKey := os.Getenv("DARAJA_PASSKEY")
	businessShortCode := int(businessShortCode64)
	password := p.GeneratePassword(businessShortCode, passKey)
	timestamp := time.Now().Format("20060102150405")
	token, err := p.GetAuthToken()
	if err != nil {
		print(err)
	}

	request := &StkPushRequest{
		BusinessShortCode: businessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            p.PhoneNumber,
		PartyB:            businessShortCode,
		PhoneNumber:       p.PhoneNumber,
		CallBackURL:       "https://mbumwa3d-production.up.railway.app/darajacallback",
		AccountReference:  "Mbumwa3D",
		TransactionDesc:   "Payment of 3D Printing",
	}

	payload, err := json.Marshal(request)

	req, err := http.NewRequest("POST", p.DarajaStkPushUrl, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := p.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
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

func DarajaCallbackHandler(w http.ResponseWriter, r *http.Request) error {
	var callbackResponse StkCallbackResponse
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

	for _, item := range stkCallback.CallbackMetadata.Item {
		fmt.Printf("%s: %v\n", item.Name, item.Value)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
