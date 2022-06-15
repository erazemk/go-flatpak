user = $(shell whoami)

all: generate permissions

generate:
	sudo go generate -run go run . -o ./pkg/

permissions:
	sudo chown -R $(user):$(user) pkg
