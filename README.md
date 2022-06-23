# ERR_CONNECTION_RESET

This service simply closes all incoming TCP connections

Why? -> because why not??

## Run locally

`PORT=4000 go run main.go`

## Run in docker

`docker run --rm  -e PORT=4000 -p 4000:4000 -it zeetdev/conn-reset`
