# Node

The base atom for the INTORCH platform

This is a Runnable part of INTORCH platform, that read messages froma channel execute an engine and put the response in another channel.

## Usage

Create new Engine

```go
var doNothingEngine = Engine(func(msg message.Message) message.Message { 
    return msg 
})
```

Create new node object and execute it

```go
nd := node.New(&nodeTestEngine)
```

Create a message and add it to the node, then run engine

```go
msg := message.New(make(message.Header), make(message.Body))
nd.Write(msg)

go func() {
    nd.Run()
}()
```

Once started the engine, the input data will be processed and you can read response.

```go
resp := nd.Read()
```

we can also use the channel to create a loop.

```go
for msg := range nd.GetReader() {
    //do something
}
```




