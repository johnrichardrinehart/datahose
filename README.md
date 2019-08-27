Datahose
=======

Datahose generates data according to a user-defined schema.

It can spray data to stdout, files, unix domain sockets, TCP/IP, and more!

# Example usage

1. Install the `datahose` CLI.
```
go get -u github.com/johnrichardrinehart/datahose/cmd/datahose
```
1. Copy the following code into a file called `data.schema` in your working directory.
```
{
	"foo": "int",
	"bar": "float32",
	"buzz": "float64",
	"baz:": "bool",
	"bizz": "string"
}
```
1. Spray data to `stdout`, a Unix domain socket (`data.unix`), a TCP/IP address (`tcp://localhost:9999`), and a file (`data.out`) by executing the following (ensure that `$GOPATH/bin` is in your path):
```
datahose -schema ./data.schema -sinks stdout,dataa.unix,tcp://localhost:9999,data.out
```

# Limitations
Currently, the schema must be flat and is constrained to Go primitive types.

# Future Work
1. Support composite types
1. Support different output wire formats (JSON, avro, MessagePack, protobuf, etc.)
1. Support workers/goroutines for increased throughput
1. Support recursive schema parsing
1. Support different schema input formats
1. Support outputting to queues (AWS SQS, Cloud PubSub, etc.)
1. Support more protocols (MQTT, HTTP, etc.)
