FROM golang:latest

LABEL maintainer="BenjaminHuang <Benjamin@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8964

RUN go build

CMD ["./srever"]