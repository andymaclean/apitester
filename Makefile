

.PHONY:  all

all:  apitester

apitester:
	go build -o bin/apitester src/main.go


