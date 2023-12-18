FROM golang:1.20

WORKDIR /app

copy go.mod .
copy controllers /app/controllers
copy models /app/models
copy main.go .

RUN go build -o go-bin .

EXPOSE 8030

ENTRYPOINT [ "/app/go-bin"]