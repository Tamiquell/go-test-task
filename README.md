### User API

##### Usage
 - Clone the repository
 - initialize module
 - use go mod tidy
 - install swag via `go get -u github.com/swaggo/swag/cmd/swag`
 - add `export PATH=$PATH:$(go env GOPATH)/bin`
 - run `go run main.go`

 The API will start on: `http://127.0.0.1:8080/api/v1/`
 Documentation is in [swagger.yaml](swagger.yaml)