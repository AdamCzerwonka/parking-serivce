FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -v -o /server cmd/server/main.go

EXPOSE 8080

CMD [ "/server"]


