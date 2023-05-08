FROM golang:1.20.4-alpine3.17

WORKDIR /game-store/api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .air.toml ./

RUN go mod download

ENV TERM=xterm-256color

CMD [ "air", "-c", ".air.toml" ]
