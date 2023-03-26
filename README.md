
# ennbu - .env Editor CLI

`ennbu` is a command-line interface (CLI) tool for managing .env files in your projects, written in Go. It allows you to easily set, replace, list, and get values of keys within your .env files.

## Features

- Get a specific key-value pair from a .env file
- List all key-value pairs in a .env file
- Set a new value for a key in a .env file, creating the key if it doesn't exist
- Replace a value for an existing key in a .env file

## Installation

There are several ways to install ennbu:

1. Using Go
To install ennbu with Go, make sure you have Go installed on your system, and then run:

```bash
go install github.com/u-yas/ennbu@latest
```

2. Downloading [from GitHub Releases](https://github.com/u-yas/ennbu/releases)

```bash
curl -L -o ennbu https://github.com/u-yas/ennbu/releases/download/v0.0.1/ennbu-linux-x86_64.tar.gz
sudo mv ennbu /usr/local/bin/
```

3. Using Docker

```bash
docker run -it -v "$(pwd):/workdir" ghcr.io/u-yas/ennbu:latest <command>
```

4. Github Action

```yml
---
name: "golang testing"
on:
  pull_request:
    branches:
      - main
jobs:
  build:
    timeout-minutes: 15
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version-file: "./go.mod"
      - name: go mod download
        run: go mod download
      - name: replace APP_ENV
        uses: u-yas/ennbu@v0.0.1
        with:
          envPath: .env
          commands: |
            set -k APP_ENV test
            list --json
      - name: go test
        run: go test ./...

```

### Usage

#### Get

To get the value of a specific key in a .env file, use the `get` command:

```bash
ennbu get -k KEY_NAME -e .env
```

#### List

To list all key-value pairs in a .env file, use the `list` command:

```bash
ennbu list -e .env
```

To list the key-value pairs in JSON format, use the `--json` flag:

```bash
ennbu list --json -e .env
```

#### Set

To set a new value for a key in a .env file, use the `set` command:

```bash
ennbu set -e .env -k KEY_NAME VALUE
```

#### Replace

To replace a value for an existing key in a .env file, use the `replace` command:

```bash
echo NEW_VALUE | ennbu replace  -e .env -k KEY_NAME
```

## Contributing

Contributions are welcome! Please feel free to submit a [pull request](https://github.com/u-yas/ennbu/pulls) or [open an issue.](https://github.com/u-yas/ennbu/issues/new)
