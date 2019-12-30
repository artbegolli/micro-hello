# micro-hello
A hello world of go-micro to demonstrate the utility of using process plugins for the ocibuilder

## Generation

go-micro uses [protoc-gen-micro](https://github.com/micro/protoc-gen-micro) to generate a `pb.micro.go` file from a `.proto` file. 

This includes clients and ahndlers which reduce boiler plate code.

You can generate by running

```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```

## Running

We can first run the server

```
go run main.go
```

And the client based on the flag

```
go run main.go --run_client
```
