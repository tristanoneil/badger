# Badger

Badger is a simplified self hosted Gist alternative built with Go.

## Setup

To get started developing locally clone down the repo then run `./setup.sh`.
This will create a local development database and add some seed data, including
a default user of:

```
user@example.com
password
```

You may retrieve Badger's dependencies using `go get ./...` then start the
server with `go install github.com/tristanoneil/badger && badger`.
