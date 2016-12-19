.PHONY: deps test build

all: test build

test:
		go test -v $$(go list ./tests/... | grep -v '/vendor/')

build:
		go build -o obd2-data-api

run: build
		./obd2-data-api

clean:
		rm -rf bin obd2-data-api
