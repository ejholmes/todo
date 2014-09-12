.PHONY: cmd

cmd:
	godep go build -o build/server ./cmd/server
