package main

import (
	"client-server-api/types"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

func main() {

	http.HandleFunc("/cotacao", BuscaCotacaoDolar)
	http.ListenAndServe(":8080", nil)
}

func BuscaCotacaoDolar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS cotacoes (
        id TEXT PRIMARY KEY,
        valor REAL NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

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

	var c types.Cotacao

	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	cotacao := newCotacao(c.USDBRL.Bid)
	err = insertCotacao(db, cotacao, ctx)
	if err != nil {
		panic(err)
	}

	cotacoes, err := selectAllCotacoes(db)
	if err != nil {
		panic(err)
	}
	for _, c := range cotacoes {
		fmt.Printf("Cotação: %v\n", c.Valor)
	}

	bid := types.CotacaoReturn{Bid: c.USDBRL.Bid}
	err = json.NewEncoder(w).Encode(&bid)
	if err != nil {
		panic(err)
	}
}

func newCotacao(valor string) *types.CotacaoDb {

	bidValue, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		panic(err)
	}

	return &types.CotacaoDb{
		ID:    uuid.New().String(),
		Valor: bidValue,
	}
}

func insertCotacao(db *sql.DB, cotacao *types.CotacaoDb, ctx context.Context) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes(id, valor) VALUES($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cotacao.ID, cotacao.Valor)
	if err != nil {
		return err
	}

	return nil
}

func selectAllCotacoes(db *sql.DB) ([]types.CotacaoDb, error) {
	rows, err := db.Query("SELECT id, valor FROM cotacoes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cotacoes []types.CotacaoDb
	for rows.Next() {
		var c types.CotacaoDb
		err = rows.Scan(&c.ID, &c.Valor)
		if err != nil {
			return nil, err
		}
		cotacoes = append(cotacoes, c)
	}
	return cotacoes, nil
}
