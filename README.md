# Badger

[![Circle CI](https://circleci.com/gh/tristanoneil/badger.svg?style=svg)](https://circleci.com/gh/tristanoneil/badger)

Badger is a simplified self hosted Gist alternative built with Go.

![Badger](http://cl.ly/image/06243x3A0i24/download/badger.gif)

## Setup

First you must create a PostgreSQL database for Badger like so: `createdb badger`.

You may retrieve Badger's dependencies using `go get ./...` then start the
server with `go install github.com/tristanoneil/badger && badger`. This
will automatically migrate the database as well.
