FROM golang:1.20-alpine 

WORKDIR /build
ADD . /build

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download

RUN GOARCH=amd64 GOOS=linux go build -o main

EXPOSE 8080

CMD ["/build/main"]