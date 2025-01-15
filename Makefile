run:
	@air -c .air-producer.toml & air -c .air-consumer.toml


load-test:
	@go run requester/main.go

PHONY: run load-test