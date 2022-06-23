FROM golang:alpine as builder

WORKDIR /app 

COPY main.go .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" main.go

FROM alpine

COPY --from=builder  /app/main /main

ENTRYPOINT ["/main"]