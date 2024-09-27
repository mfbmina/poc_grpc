# PoC gRPC

Used as a PoC for:
- [Blog post PT-BR](https://mfbmina.dev/posts/golang-grpc/)
- [Blog post EN](https://mfbmina.dev/en/posts/golang-grpc/)

It defines a gRPC server and a client.

The server will receive a request from the client and respond with PONG if the message is PING.

The client will wait for an input from the user and send it to the server.

## Running

1. `$ go run client/main.go`
1. `$ go run server/main.go`
