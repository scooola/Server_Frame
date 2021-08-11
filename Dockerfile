# build executable program
FROM golang:alpine AS builder
WORKDIR /usr/src
COPY . .
RUN go build -v -mod=vendor -o server

# run server program
FROM alpine:latest
WORKDIR /usr/local/bin
# copy program
COPY --from=builder /usr/src/server .
# copy server config
COPY --from=builder /usr/src/server_config.json .
CMD ["./server", " --port 12138", " --gin_mode release"]
