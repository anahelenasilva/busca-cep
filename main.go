package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anahelenasilva/busca-cep/entities"
)

func main() {

	file, err := os.Create("cep_data.json")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	var cepsData []entities.Cep

	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", url, "-", err.Error())
			continue
		}

		defer req.Body.Close()

		response, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error reading response body for URL:", url, "-", err.Error())
			continue
		}

		var cep entities.Cep
		err = json.Unmarshal(response, &cep)
		if err != nil {
			fmt.Println("Error unmarshalling JSON for URL:", url, "-", err.Error())
			continue
		}

		cepsData = append(cepsData, cep)
	}

	jsonData, err := json.MarshalIndent(cepsData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON data:", err.Error())
		os.Exit(1)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err.Error())
		os.Exit(1)
	}
}
