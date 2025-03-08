package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
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

func main() {
	http.HandleFunc("/", buscaCotacaoDolar)
	http.ListenAndServe(":8080", nil)
}

func buscaCotacaoDolar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var c Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(c.USDBRL.Bid))

	log.Println("Request finalizada com sucesso")
}
