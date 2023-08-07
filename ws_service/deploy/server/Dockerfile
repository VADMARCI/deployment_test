FROM golang:alpine
ENV CGO_ENABLED=0
RUN apk update && apk add bash inotify-tools git curl nano

WORKDIR /ws_service

COPY ws_service/go.* ./
COPY core/go.* ../core/
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY core  ../core
COPY ws_service ./

RUN go build -gcflags "all=-N -l" -o ./bin/server ./cmd/server/main.go

