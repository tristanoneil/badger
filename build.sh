GOOS=linux GOARCH=386 go build -o builds/badger.linux
GOOS=darwin GOARCH=386 go build -o builds/badger.mac
rice -i ./models -i ./routes append --exec builds/badger.linux
rice -i ./models -i ./routes append --exec builds/badger.mac
