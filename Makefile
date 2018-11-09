all:
	go build ./src/app	

build_linux:
	env GOOS=linux GOARCH=amd64 go build ./src/app
