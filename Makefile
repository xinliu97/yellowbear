.PHONY: run, sample

run:
	go run ~/yellowbear/cmd/main.go

create_sample:
	go run ~/yellowbear/cmd/main.go -sample