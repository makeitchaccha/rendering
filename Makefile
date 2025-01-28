fmt:
	go fmt ./...
.PHONY: fmt

lint:
	staticcheck ./...
.PHONY: lint

vet:
	go vet ./...
.PHONY: vet
