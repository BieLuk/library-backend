FROM golang:1.20

WORKDIR /app
ADD . /app
RUN go build -o /library-backend ./cmd/web/*
EXPOSE 8080

CMD ["/library-backend"]