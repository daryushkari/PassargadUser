# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 9000
EXPOSE 8080

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /PassargadUser
CMD ["/PassargadUser"]