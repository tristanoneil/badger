# Badger

Badger is a simplified self hosted Gist alternative built with Go.

## Setup

First you must create a PostgreSQL database for Badger like so: `createdb badger`.

You may retrieve Badger's dependencies using `go get ./...` then start the
server with `go install github.com/tristanoneil/badger && badger`. This
will automatically migrate the database as well.
