package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (p *PaymentProcessor) InitiateStkPush() {
	businessShortCode := 174379
	passKey := "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
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
		Amount:            20,
		PartyA:            p.PhoneNumber,
		PartyB:            businessShortCode,
		PhoneNumber:       p.PhoneNumber,
		CallBackURL:       "https://mydomain.com/path",
		AccountReference:  "CompanyXLTD",
		TransactionDesc:   "Payment of 3D Printing at mbumwa3D",
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
