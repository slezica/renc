build:
	mkdir -p bin
	go build -o bin/renc src/*.go

build-all:
	mkdir -p bin

	env GOOS=linux GOARCH=386 go build -o bin/renc-linux32 src/*.go
	env GOOS=windows GOARCH=386 go build -o bin/renc-windows32.exe src/*.go

	env GOOS=darwin GOARCH=amd64 go build -o bin/renc-darwin64 src/*.go
	env GOOS=linux GOARCH=amd64 go build -o bin/renc-linux64 src/*.go
	env GOOS=windows GOARCH=amd64 go build -o bin/renc-windows64.exe src/*.go
