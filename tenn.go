package tenn

import (
	"net/http"
	"io/ioutil"
)

type Tenn struct {
	BaseURL string
	ApiKey string
	DepositURL string
	WithdrawURL string
}

func (t *Tenn) New(apiKey string, depositURL string, withdrawURL string) {
	t.BaseURL = "https://weprocesspayments.ink/graphql"
	t.ApiKey = apiKey
	t.DepositURL = depositURL
	t.WithdrawURL = withdrawURL
}

//initiate deposit
/*
mutation {
  initateDeposit(input:{
    amountInCents: 1000
    userPhoneNumber: "254112159579"
  })
  {
    Txid
    Status
  }
}
*/
func (t *Tenn) InitiateDeposit(amountInCents int, userPhoneNumber string) (string, error) {
	//send request
	req, err := http.NewRequest("POST", t.BaseURL, nil)
	if err != nil {
		return "", err
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + t.ApiKey)

	//set body
	body := `{"query":"mutation {initateDeposit(input:{amountInCents: ` + string(amountInCents) + `, userPhoneNumber: \"` + userPhoneNumber + `\"}){Txid Status}}","variables":null}`
	req.Body = body

	//send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	//read response
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

//initiate withdrawal
/*
mutation {
  initateWithdrawal(input: {
    amountInCents: 1001,
    userPhoneNumber: "254112159579"
  }){
    Txid,
    Status
    Message
  }
}
*/

func (t *Tenn) InitiateWithdrawal(amountInCents int, userPhoneNumber string) (string, error) {
	//send request
	req, err := http.NewRequest("POST", t.BaseURL, nil)
	if err != nil {
		return "", err
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + t.ApiKey)

	//set body
	body := `{"query":"mutation {initateWithdrawal(input: {amountInCents: ` + string(amountInCents) + `, userPhoneNumber: \"` + userPhoneNumber + `\"}){Txid, Status, Message}}","variables":null}`
	req.Body = body

	//send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	//read response
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
