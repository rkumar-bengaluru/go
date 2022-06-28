# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app 
COPY go.mod ./
COPY go.sum ./
COPY ./cmd ./cmd
COPY ./handler ./handler
COPY ./model ./model
COPY ./resources ./resources
COPY ./server ./server
RUN go build -o /docker-go-server ./cmd
EXPOSE 8080
CMD [ "/docker-go-server" ]