FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o server .


FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/server .
ENV PORT=8080
EXPOSE 8080
CMD ["./server"]

