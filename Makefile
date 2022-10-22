# port for main server
port := 5053
moc_server_address := 127.0.0.1:5059
bold := \033[1m
normal := \033[0m
good := \033[1m\033[0;32m

help:
	@echo "$(bold)Makefile commands$(normal)"
	@echo "-----------------"
	@echo "$(bold)make build$(normal)   : will build the project"
	@echo "$(bold)make tests$(normal)   : run tests for the project"
	@echo "$(bold)make run$(normal)     : will run the project"
	@echo ""
	@echo "$(bold)OS commands$(normal)"
	@echo "-----------"
	@echo "start server at PORT with 'IP:PORT' list of partners:"
	@echo "$(bold)./bin/simple-choose-ad -p PORT -d 'IP:PORT'$(normal)"

run:
	go run cmd/main.go -p $(port) -d "$(moc_server_address)"

test-ip:
	@echo
	@go run cmd/main.go -p $(port) -d "$(moc_server_address),localhost:5059" || \
	{ echo "\n[+] PASS wrong IP address test"; exit 0; }

test-port:
	@echo
	@go run cmd/main.go -p $(port) -d "$(moc_server_address),127.0.0.1:as" || \
	{ echo "\n[+] PASS wrong port test"; exit 0; }

test-port-max:
	@echo
	@go run cmd/main.go -p $(port) -d "$(moc_server_address),127.0.0.1:65537" || \
	{ echo "\n[+] PASS port too big test"; exit 0; }

test-port-endpoint:
	@echo
	@go run cmd/main.go -p $(port) -d "127.0.0.1:9001/bid_request" || \
	{ echo "\n[+] PASS endpoint with address test"; exit 0; }

build:
	go build -o bin/simple-choose-ad cmd/main.go

build-and-push:
	@GOOS=linux GOARCH=amd64 go build -o build/ssp cmd/main.go && rsync -ah build/ssp ubuntu:~/ssp-testbed-clone/builds

start-moc-server:
	@echo "[!] Starting up moc-server on $(moc_server_address) ..."
	@go run internal/moc_server.go -l $(moc_server_address) &


stop-moc-server:
	@echo "[!] Stopping moc-server ..."
	@curl -s -o /dev/null "$(moc_server_address)/exit" &

test-server:
	@echo
	@echo "Check response from moc-server "
	@$(MAKE) start-moc-server
	@cd "cmd/client_server/"; \
	go test -v
	@$(MAKE) stop-moc-server

tests:
	# @$(MAKE) test-ip
	# @$(MAKE) test-port
	# @$(MAKE) test-port-max
	# @$(MAKE) test-port-endpoint
	@$(MAKE) test-server
