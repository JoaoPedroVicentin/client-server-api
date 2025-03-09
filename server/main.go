package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type USDBRL struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Cotacao struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type CotacaoReturn struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", BuscaCotacaoDolar)
	http.ListenAndServe(":8080", nil)
}

func BuscaCotacaoDolar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var c Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	bid := CotacaoReturn{Bid: c.USDBRL.Bid}
	err = json.NewEncoder(w).Encode(&bid)
	if err != nil {
		panic(err)
	}
}
