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

		var data entities.Cep
		err = json.Unmarshal(response, &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON for URL:", url, "-", err.Error())
			break
		}

		fmt.Printf("CEP data: %s\n", data)
	}
}
