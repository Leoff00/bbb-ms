run:
	@air -c .air-producer.toml & air -c .air-consumer.toml

run-no-air:
	go run cmd/main.go & go run consumer/consumer_main.go

load-test:
	@go run requester/main.go

PHONY: run load-test run-no-air