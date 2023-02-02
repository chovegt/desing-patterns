# Desing patterns

This project is for a practices, and enjoy :D

## How to Install and Run the Project
requires [Go](https://go.dev/doc/install) v1.19+ to run.

Install the dependencies and build app.

```sh
go mod init
go mod tidy

go build ./...
```

Environments...

| ENV | VALUE |
| ------ | ------ |
| SCOPE | local |
| XXXX | xxxx |


### vscode
create launch.json

```json
   "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env":{
                "SCOPE":"local"
            }
        }
    ]
```

## API endpoints
[WIP]

## Authors
[Diogenes Salazar](@DioCoding)