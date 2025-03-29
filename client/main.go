package main

import (
	"client-server-api/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	req, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
	}
	var data types.CotacaoReturn
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
