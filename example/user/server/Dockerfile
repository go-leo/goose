FROM golang:1.23

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o user ./main.go

EXPOSE 8000

ENTRYPOINT [ "/app/user" ]
