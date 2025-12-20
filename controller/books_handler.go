package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui no Search")
	
	if r.Body != nil {
		defer r.Body.Close()
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, "Erro ao ler corpo da requisição")
		return
	}

	query := strings.TrimSpace(string(b))
	fmt.Println("Buscando por:", query)

	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gReq, err := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}