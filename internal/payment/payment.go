package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

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
	req.Header.Add("Host", "sandbox.safaricom.co.ke")

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

func (p *PaymentProcessor) InitiateStkPush(amount int) (*TransactionResponse, error) {
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
		CallBackURL:       "https://3d.mbumwa.com/darajacallback",
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
	if err != nil {
		return nil, err
	}

	var transactionResponse TransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return nil, err
	}

	if transactionResponse.ResponseCode != "0" {
		return nil, errors.New(transactionResponse.CustomerMessage)
	}

	return &transactionResponse, nil
}
