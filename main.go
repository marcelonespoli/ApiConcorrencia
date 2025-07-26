package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	Api     string
	Content []byte
}

func RequestAPI(ctx context.Context, url string, apiName string, ch chan<- Result) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao iniciar new request com context")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao realizar request a API " + apiName)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler response da API " + apiName)
	}

	var response Result
	response.Api = apiName
	response.Content = body
	ch <- response
}

func main() {
	cep := "01153000"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultChan := make(chan Result, 2)

	brasilAPI := "https://brasilapi.com.br/api/cep/v1/" + cep
	viaCepAPI := "http://viacep.com.br/ws/" + cep + "/json/"

	go RequestAPI(ctx, viaCepAPI, "ViaCEP", resultChan)
	go RequestAPI(ctx, brasilAPI, "BrasilAPI", resultChan)

	select {
	case result := <-resultChan:
		fmt.Println("\n=== API Vencedora:", result.Api, "===")
		fmt.Printf("API Response: \n%s\n", string(result.Content))
	case <-ctx.Done():
		fmt.Println("Erro: timeout ao buscar dados das APIs")
	}
}
