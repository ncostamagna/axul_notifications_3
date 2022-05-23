.PHONY: build clean deploy
export GOSUMDB=off

install:
	git config --local core.hooksPath .githooks/
	export GOSUMDB=off
	npm i
	git config url."git@gitlab.com:".insteadOf "https://gitlab.com/"
	go mod tidy
	go mod download

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/notification-getall 		    cmd/notification/getall/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/notification-create		    	cmd/notification/getall/main.go

start:
	make build
	sls offline --printOutput

dev:
	docker-compose up -d
	make build
	sls offline --printOutput

deploy:
	make build
	sls deploy

test-cover:
	go test ./internal/... ./pkg/... -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

test:
	go test ./internal/... ./pkg/... -v

