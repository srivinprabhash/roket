default: run

run:
	go run main.go

build:
	rm -rf build/rocket
	go build -o build/rocket -ldflags "-s -w" main.go