FROM golang:1.22-rc-bookworm

WORKDIR app

COPY . ./

RUN go mod download

RUN go build -o main .

EXPOSE 3001

CMD ["./main"]