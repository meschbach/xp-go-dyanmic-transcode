#  gRPC Protobuf to JSON service
An experiment to determine the technique to translate an arbitrary protobuf message into JSON on the fly.

## Checking it out
To build just run `./build.sh`.  To see it in action:
```shell
./service/service &
./client/client
```

[service/transcoder.go](service/transcoder.go) contains the key features within the service side to reflect on the
incoming type.

[client/main.go](client/main.go) starting around `OtherMessage` show how to produce the workable generic handling.

## References
* [Technique](https://protobuf.dev/programming-guides/techniques/)
* https://protobuf.dev/programming-guides/techniques/#self-description
* https://github.com/golang/protobuf/issues/1215#issuecomment-706593142

### `protoreflect`
* https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect
### `dyanmicpb` attempts
* [godoc](https://pkg.go.dev/google.golang.org/protobuf@v1.33.0/types/dynamicpb#NewMessageType)
* [Go source](https://go.googlesource.com/protobuf/+/refs/heads/master/types/dynamicpb/dynamic.go?autodive=0%2F%2F%2F%2F%2F)
* [RPC UNARY REQUESTS THE HARD WAY: USING PROTOREFELECT, DYNAMICPB AND WIRE-ENCODING TO SEND MESSAGES](https://blog.salrashid.dev/articles/2022/grpc_wireformat/)
* [`grpc_wireformat`](https://github.com/salrashid123/grpc_wireformat/blob/main/grpc_client_dynamic.go)
* [SO using `dyanmicpb`](https://stackoverflow.com/questions/65242456/convert-protobuf-serialized-messages-to-json-without-precompiling-go-code)
