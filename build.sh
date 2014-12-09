( cd models && rice embed-go )
( cd routes && rice embed-go )
GOOS=linux GOARCH=386 go build -o builds/badger.linux
