# Busca CEP

A simple Go program to fetch and display Brazilian postal code (CEP) information from web APIs.

## Description

This program accepts one or more URLs as command-line arguments, fetches data from each URL, and parses the response as CEP information. It's designed to work with Brazilian postal code APIs that return JSON data with address details.

## Features

- Fetch CEP data from multiple URLs in a single execution
- Parse JSON responses containing Brazilian address information
- Display structured CEP data including:
  - CEP (postal code)
  - Logradouro (street address)
  - Complemento (complement)
  - Unidade (unit)
  - Bairro (neighborhood)
  - Localidade (city)
  - UF (state abbreviation)
  - Estado (state name)
  - Região (region)
  - IBGE code
  - GIA code
  - DDD (area code)
  - SIAFI code

## Prerequisites

- Go 1.16 or later installed on your system

## Installation

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd busca-cep
   ```

## Usage

### Basic Usage

```bash
go run main.go <URL1> <URL2> ... <URLn>
```

### Examples

#### Using ViaCEP API (most common Brazilian CEP API)

```bash
# Fetch information for a single CEP
go run main.go "https://viacep.com.br/ws/01310-100/json/"

# Fetch information for multiple CEPs
go run main.go "https://viacep.com.br/ws/01310-100/json/" "https://viacep.com.br/ws/20040-020/json/"
```

#### Example Output

```
CEP data: {01310-100 Avenida Paulista  {} Bela Vista São Paulo SP São Paulo Sudeste 3550308  11 1004}
CEP data: {20040-020 Avenida Rio Branco de 1 ao 185 - lado ímpar {} Centro Rio de Janeiro RJ Rio de Janeiro Sudeste 3304557  21 6001}
```

### Building the Program

To create an executable binary:

```bash
go build -o busca-cep main.go
```

Then run the binary:

```bash
./busca-cep "https://viacep.com.br/ws/01310-100/json/"
```

## API Compatibility

This program is designed to work with APIs that return JSON data in the following format:

```json
{
  "cep": "01310-100",
  "logradouro": "Avenida Paulista",
  "complemento": "",
  "unidade": "",
  "bairro": "Bela Vista",
  "localidade": "São Paulo",
  "uf": "SP",
  "estado": "São Paulo",
  "regiao": "Sudeste",
  "ibge": "3550308",
  "gia": "",
  "ddd": "11",
  "siafi": "1004"
}
```

## Common CEP APIs

- **ViaCEP**: `https://viacep.com.br/ws/{CEP}/json/`
- **CEP Aberto**: `https://www.cepaberto.com/api/v3/cep?cep={CEP}` (requires API key)
- **PostalMon**: `https://api.postmon.com.br/v1/cep/{CEP}`

## Error Handling

The program handles the following error scenarios:

- Network connection errors when fetching URLs
- Invalid or malformed JSON responses
- HTTP request failures

If an error occurs, the program will display an error message and stop processing further URLs.

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

This project is open source and available under the [MIT License](LICENSE).