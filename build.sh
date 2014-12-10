GOOS=linux GOARCH=386 go build -o builds/badger.linux
rice -i ./models -i ./routes append --exec builds/badger.linux
