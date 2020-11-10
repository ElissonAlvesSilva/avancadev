FROM golang:1.15

WORKDIR /go/src/fullcycle

COPY . .

RUN go mod download

RUN GOOS=linux go build -o ./fullcycle .

ENTRYPOINT ["./fullcycle"]