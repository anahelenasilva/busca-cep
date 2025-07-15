# Busca CEP

A simple Go web server that provides a REST API to fetch Brazilian postal code (CEP) information.

## Description

This program runs as an HTTP server that exposes a REST API endpoint to fetch CEP (Brazilian postal code) information. The server integrates with the ViaCEP API to retrieve detailed address information and returns it as JSON responses to client requests.

## Features

- HTTP REST API server for CEP lookups
- Integration with ViaCEP API for reliable data
- JSON response format
- Error handling with proper HTTP status codes
- Simple query parameter-based requests
- Returns comprehensive address information including:
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
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on port 8090 and display the message "Server running on port 8090".

## Usage

### Starting the Server

```bash
go run main.go
```

### API Endpoint

The server exposes a single endpoint:

**GET** `/busca-cep?cep={CEP_CODE}`

### Examples

#### Fetch CEP information

```bash
# Using curl
curl "http://localhost:8090/busca-cep?cep=01310-100"

# Using a web browser
http://localhost:8090/busca-cep?cep=01310-100
```

#### Multiple requests

```bash
# Fetch different CEPs
curl "http://localhost:8090/busca-cep?cep=01310-100"
curl "http://localhost:8090/busca-cep?cep=20040-020"
curl "http://localhost:8090/busca-cep?cep=30112-000"
```

#### Example Response

**Successful Request:**
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

**Error Response (Missing CEP parameter):**
```
HTTP 400 Bad Request
CEP parameter is required
```

**Error Response (Invalid CEP or network error):**
```
HTTP 500 Internal Server Error
Error message details
```

### Building the Program

To create an executable binary:

```bash
go build -o busca-cep main.go
```

Then run the binary:

```bash
./busca-cep
```

## API Reference

### Endpoint

**GET** `/busca-cep`

### Parameters

| Parameter | Type   | Required | Description                    |
|-----------|--------|----------|--------------------------------|
| `cep`     | string | Yes      | Brazilian postal code to search |

### Response Format

The API returns JSON data in the following format:

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

### HTTP Status Codes

| Status Code | Description                    |
|-------------|--------------------------------|
| 200         | Success - CEP found            |
| 400         | Bad Request - Missing CEP parameter |
| 500         | Internal Server Error - API or network error |

## Integration

This server integrates with the **ViaCEP API** (`https://viacep.com.br/ws/{CEP}/json/`) to fetch reliable CEP information. ViaCEP is a free and widely-used Brazilian postal code API.

## Error Handling

The server handles the following error scenarios:

- **Missing CEP parameter**: Returns HTTP 400 with error message
- **Network connection errors**: Returns HTTP 500 with error details
- **Invalid or malformed JSON responses**: Returns HTTP 500 with error details
- **ViaCEP API failures**: Returns HTTP 500 with error details

All errors include descriptive messages to help with debugging.

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

This project is open source and available under the [MIT License](LICENSE).