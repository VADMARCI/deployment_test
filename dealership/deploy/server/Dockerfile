FROM golang:alpine
ENV CGO_ENABLED=0
RUN apk update && apk add bash inotify-tools git curl nano

WORKDIR /dealership

COPY dealership/go.* ./
COPY core/go.* ../core/
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY core  ../core
COPY dealership ./


RUN go build -gcflags "all=-N -l" -o ./bin/server ./cmd/server/main.go

