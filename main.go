package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", url, "-", err.Error())
			break
		}

		defer req.Body.Close()

		response, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading response body for URL:", url, "-", err.Error())
			break
		}

		var data Cep
		err = json.Unmarshal(response, &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON for URL:", url, "-", err.Error())
			break
		}

		fmt.Printf("CEP data: %s\n", data)
	}
}
