go-bindata -o static/bindata.go dist/... migrations/... templates/...
sed -i '' 's/main/static/' static/bindata.go

GOOS=linux GOARCH=386 go build -o builds/badger.linux
GOOS=darwin GOARCH=386 go build -o builds/badger.mac
