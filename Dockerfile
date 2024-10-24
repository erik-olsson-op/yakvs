FROM golang:1.23

WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o api

EXPOSE 8443
EXPOSE 7443

CMD ["./api"]