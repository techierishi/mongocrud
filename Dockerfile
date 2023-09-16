FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mongocrud

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mongocrud .

EXPOSE 8060

CMD ["./mongocrud"]
