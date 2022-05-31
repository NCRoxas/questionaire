run:
	go run main.go

clean:
	rm questionaire

compile:
	echo "Compiling..."
	GOOS=linux GOARCH=amd64 go build -o questionaire main.go
