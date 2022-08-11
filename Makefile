default: run

run:
	go run main.go

clean:
	@rm -rf build/

build:
	@go build -o build/roket -ldflags "-s -w" main.go