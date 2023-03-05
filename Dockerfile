FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./out/dist .

RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf

CMD ./out/dist
