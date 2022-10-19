# build stage
FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ARG GOPROXY
ENV GOPROXY=${GOPROXY}

WORKDIR $GOPATH/src/gitlab.com/ht-co/doudou/doudou/ambipom

COPY go.mod go.sum ./

RUN apk add --no-cache git && \
    git config --global url."https://gitlab+deploy-token-1270183:n8DF2s8Ap864-zCLxkcN@gitlab.com/ht-co/wraperr".insteadOf "https://gitlab.com/ht-co/wraperr" && \
    git config --global url."https://gitlab+deploy-token-1397536:KWRg9Xmk-nksLWHzPxFC@gitlab.com/ht-co/doudou/live/enum-platform-go".insteadOf "https://gitlab.com/ht-co/doudou/live/enum-platform-go" && \
    go mod download

COPY . .

RUN go build -ldflags '-s -w' -o /app ./cmd/app/main.go

# final stage
FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /app .
COPY config/app.conf.example.yaml ./app.conf.yaml

ENTRYPOINT ["./app"]
