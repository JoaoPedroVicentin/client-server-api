package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	// defer cancel()
	// req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	req, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
	}
	var data CotacaoReturn
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao tentar fazer parse da resposta: %v\n", err)
	}
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
	defer file.Close()
	_, _ = file.WriteString(fmt.Sprintf("Dólar: %s", data.Bid))
}
