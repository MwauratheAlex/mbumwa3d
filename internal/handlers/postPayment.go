package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
)

func PostPayment(w http.ResponseWriter, r *http.Request) error {
	_, ok := r.Context().Value(middleware.UserKey).(*store.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, components.UnauthorizedFormEror())
	}

	authToken, err := GetAuthToken()
	fmt.Println("AUTH TOKEN: ", authToken, err)
	phoneNumber := r.FormValue("phone")

	fmt.Println(phoneNumber)
	return nil
}

func GetAuthToken() (string, error) {
	url := os.Getenv("DARAJA_AUTH_URL")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	darajaAuthHeader := os.Getenv("DARAJA_AUTH_HEADER")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", darajaAuthHeader)

	res, err := client.Do(req)
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

type StkPushRequest struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int    `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

func InitiateStkPush() {
	url := os.Getenv("DARAJA_STK_PUSH_URL")
	request := &StkPushRequest{
		BusinessShortCode: "174379",
		Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjQwNjI5MTMwNjEw",
		Timestamp:         "20240629130610",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            1,
		PartyA:            "254708374149",
		PartyB:            "174379",
		PhoneNumber:       "254708374149",
		CallBackURL:       "https://mydomain.com/path",
		AccountReference:  "CompanyXLTD",
		TransactionDesc:   "Payment of X",
	}
	client := &http.Client{}

	payload, err := json.Marshal(request)

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 7fOKzRQAfprgIp1Ka5rL1FCTdSfe")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
