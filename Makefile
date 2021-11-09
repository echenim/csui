
test:
	go test ./... -race -cover -v

demo:
	go run examples/columns/main.go
