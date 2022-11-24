run:
	go run main.go

# go install github.com/swaggo/swag/cmd/swag@latest
generate-api-docs:
	swag init
