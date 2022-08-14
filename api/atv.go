package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
)

type Bot struct {
    Message string  `json:"message"`
    Serverless  int32  `json:"status"`
    Account *alpaca.Account `json:"account"`
}

func Atv(w http.ResponseWriter, r *http.Request) {
    // initialize alpaca client
    // uses env and tested on paper-api
    exchange := alpaca.NewClient(alpaca.ClientOpts{
        BaseURL:   "https://paper-api.alpaca.markets",
    })
    acct, err := exchange.GetAccount()
    if err != nil {
        panic(err)
    }
    bot := Bot{
        Message: "Hi from xl",
        Serverless: http.StatusOK,
        Account: acct,
    }
    fmt.Printf("%+v\n", bot)
    // Displays type of client resp
    // fmt.Printf("%T\n", acct)
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(bot)
	if err != nil {
		fmt.Println("Error in JSON marshall.\n%s", err)
	} else {
		w.Write(jsonResp)
	}
	return
}