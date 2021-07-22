all: build

build:
	go build -o ftp-server ./cmd/ftp-server/main.go
