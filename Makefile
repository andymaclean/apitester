

test_service_addr:="localhost:9995"
test_service_url:="http://$(test_service_addr)"

.PHONY:  all

all:  bin/apitester

bin/apitester: src/apitester_main/*.go src/apitester/*.go
	go build -o bin/apitester src/apitester_main/apitester.go

bin/test_server: src/test_server/*.go src/test_server_main/*.go
	go build -o bin/test_server src/test_server_main/test_server.go

e2etest:  bin/apitester bin/test_server
	@echo "TEST SERVER ON $(test_service_addr)"	
	./scripts/test_with_server ./bin/test_server --address $(test_service_addr) -- ./bin/apitester --verbose --baseurl $(test_service_url)

