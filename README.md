# Golang Calculator API implementation

Simple API calculator implemented in Golang. For request/response scheme see [postman collection](tests/postman_collections/calculator.json)

## Running project

```bash
go run ./cmd/calc_service/
# :port - Optional parameter (default: 8080)
```

## Running tests

```bash
go test ./... -count=1  -v 
```

## Sample requests

See [postman collection](tests/postman_collections/calculator.json)
