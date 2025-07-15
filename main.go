package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/anahelenasilva/busca-cep/entities"
)

func main() {
	http.HandleFunc("/busca-cep", BuscaCepHandler)
	fmt.Println("Server running on port 8090")
	http.ListenAndServe(":8090", nil)
}

func BuscaCepHandler(writer http.ResponseWriter, request *http.Request) {
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		writer.WriteHeader(http.StatusBadRequest)
		http.Error(writer, "CEP parameter is required", http.StatusBadRequest)
		return
	}

	cepResponse, err := BuscaCep(cepParam)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cepResponse)
}

func BuscaCep(cep string) (*entities.Cep, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching CEP: %w", err)
	}

	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var cepResponse entities.Cep
	err = json.Unmarshal(response, &cepResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON from provider: %w", err)
	}

	return &cepResponse, nil
}
