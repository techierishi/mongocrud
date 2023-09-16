FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mongocrud

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mongocrud .

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
RUN apk update
RUN apk add --no-cache mongodb
RUN apk add --no-cache mongodb-tools
USER root
RUN mkdir -p /data/db/
RUN chown root /data/db

EXPOSE 7070

CMD mongod --fork --logpath /var/log/mongodb.log && ./mongocrud
