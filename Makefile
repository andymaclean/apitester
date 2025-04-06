

.PHONY:  all

all:  apitester

apitester:
	go build -o bin/apitester src/commands/apitester.go

test_server:
	go build -o bin/test_server src/commands/test_server.go

