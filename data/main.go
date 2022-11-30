package main

import (
	"context"
	"fmt"

	"github.com/babulalt/go-khalti/khalti"
)

var pin string

func main() {
	clientId := "test_public_key_b048b45bdccb43ba818968273ffd49c4"
	secretId := "test_secret_key_2f39f7070aec464c8ef55beaa63df52c"
	khaltiService, err := khalti.NewKhaltiClient(clientId, secretId, nil)
	if err != nil {
		//handle error case
	}
	trasaction := &khalti.InitiateTransactionRequest{
		PubicKey:        khaltiService.ClientID,
		Mobile:          "9863857035",
		TransactionPin:  "1234",
		Amount:          1000,
		ProductIdentity: "Test",
		ProductName:     "test",
		ProductUrl:      "",
	}
	data, err := khaltiService.InitiateTransaction(trasaction)
	if err != nil {
		fmt.Println("initiate payment error ::", err)
		return
	}
	fmt.Println("Initiate Tansaction Response")
	fmt.Println(data)
	fmt.Println("***********************************************")
	fmt.Println("Enter Confirmation Code :: ")
	fmt.Scanln(&pin)
	fmt.Println("entered pin :: ", pin)
	confirm := &khalti.ConfirmTransactionRequest{
		PubicKey:         khaltiService.ClientID,
		Token:            data.Token,
		ConfirmationCode: pin,
		TransactionPin:   "1234",
	}
	fmt.Println("confirm paylaod :: ", confirm)
	con, err := khaltiService.ConfirmationTransaction(context.Background(), confirm)
	if err != nil {
		fmt.Println("confirm payment error ::", err)
		return
	}
	fmt.Println("Confirm Tansaction Response")
	fmt.Println(con)
	fmt.Println("***********************************************")
	verify := &khalti.VerifyTransactionRequest{
		Token:  data.Token,
		Amount: 1000,
	}
	fmt.Println("verify paylaod :: ", verify)
	ver, err := khaltiService.VerifyTransaction(context.Background(), verify)
	if err != nil {
		fmt.Println("confirm payment error ::", err)
		return
	}
	fmt.Println("Verify Tansaction Response")
	fmt.Println(ver)
	fmt.Println("data", data)
}
