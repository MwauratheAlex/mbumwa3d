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

func NewPaymentProcessor() *PaymentProcessor {
	businessShortCode64, err := strconv.ParseInt(os.Getenv("DARAJA_SHORTCODE"), 10, 64)
	if err != nil {
		fmt.Println("Error creating payment processor")
		panic(err)
	}
	return &PaymentProcessor{
		darajaAuthUrl:              os.Getenv("DARAJA_AUTH_URL"),
		darajaStkPushUrl:           os.Getenv("DARAJA_STK_PUSH_URL"),
		darajaTransactionStatusUrl: os.Getenv("DARAJA_TRANSACTION_STATUS_URL"),
		client:                     &http.Client{},
		passKey:                    os.Getenv("DARAJA_PASSKEY"),
		businessShortCode:          businessShortCode64,
	}
}

func (p *PaymentProcessor) getAuthToken() (string, error) {

	req, err := http.NewRequest("GET", p.darajaAuthUrl, nil)
	if err != nil {
		return "", err
	}
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	authString := fmt.Sprintf("%s:%s", consumerKey, consumerSecret)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))
	darajaAuthHeader := fmt.Sprintf("Basic %s", encodedAuth)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", darajaAuthHeader)
	req.Header.Add("Host", "sandbox.safaricom.co.ke")

	res, err := p.client.Do(req)
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

func (p *PaymentProcessor) generatePassword(shortcode int64, passkey string) string {
	timestamp := time.Now().Format("20060102150405")
	data := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
	password := base64.StdEncoding.EncodeToString([]byte(data))
	return password
}

func (p *PaymentProcessor) InitiateStkPush(amount, phoneNumber int) (
	*TransactionResponse, error) {
	password := p.generatePassword(p.businessShortCode, p.passKey)
	timestamp := time.Now().Format("20060102150405")
	token, err := p.getAuthToken()
	if err != nil {
		print(err)
	}

	request := &StkPushRequest{
		BusinessShortCode: p.businessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            phoneNumber,
		PartyB:            p.businessShortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       "https://3d.mbumwa.com/darajacallback",
		AccountReference:  "Mbumwa3D",
		TransactionDesc:   "Payment of 3D Printing",
	}

	payload, err := json.Marshal(request)

	req, err := http.NewRequest("POST", p.darajaStkPushUrl, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := p.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
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

func (p *PaymentProcessor) GetTransactionStatus(checkoutRequestId string) (*TransactionStatusResponse, error) {
	requestData := &TransactionStatusRequest{
		Initiator:                os.Getenv("DARAJA_TRANSACTION_STATUS_INITIATOR"),
		SecurityCredential:       os.Getenv("DARAJA_TRANSACTION_STATUS_CREDENTIALS"),
		CommandID:                "TransactionStatusQuery",
		TransactionID:            checkoutRequestId,
		OriginatorConversationID: checkoutRequestId,
		PartyA:                   p.businessShortCode,
		IdentifierType:           "4", // 4 - Organization shortcode (BusinessShortCode),
		ResultURL:                "https://3d.mbumwa.com/payment-status-callback",
		QueueTimeOutURL:          "https://3d.mbumwa.com/payment-status-callback",
		Remarks:                  "OK",
		Occasion:                 "Mbumwa3d Transaction",
	}
	fmt.Println("Transaction Status Request Object")
	jsonData, _ := json.MarshalIndent(requestData, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", p.darajaTransactionStatusUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	token, err := p.getAuthToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var statusResponse TransactionStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&statusResponse); err != nil {
		return nil, err
	}

	return &statusResponse, nil
}
