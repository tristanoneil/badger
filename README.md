# Badger

[![Circle CI](https://circleci.com/gh/tristanoneil/badger.svg?style=svg)](https://circleci.com/gh/tristanoneil/badger)

Badger is a simplified self hosted Gist alternative built with Go.

![Badger](http://cl.ly/image/06243x3A0i24/download/badger.gif)

## Development Setup

1. Create a PostgreSQL database for Badger like so: `createdb badger`.
1. Retrieve Badger's dependencies `go get ./...`
1. Install godotenv and fresh like so.

    ```
    $ go get github.com/pilu/fresh
    $ go get github.com/joho/godotenv/cmd/godotenv
    ```

1. Copy `.env.example` to `.env` and configure the environment
variables as necessary.

1. Start the application with `godotenv fresh`, fresh will automatically reload
your code as you make changes.
