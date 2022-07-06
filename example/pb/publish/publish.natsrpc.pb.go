// Code generated by protoc-gen-natsrpc. DO NOT EDIT.
// versions:
// - protoc             v3.14.0
// - protoc-gen-natsrpc v0.5.0
// source: publish.proto

package publish

import (
	context "context"
	fmt "fmt"
	natsrpc "github.com/byebyebruce/natsrpc"
	testdata "github.com/byebyebruce/natsrpc/testdata"
	proto "github.com/golang/protobuf/proto"
	nats_go "github.com/nats-io/nats.go"
)

var _ = new(context.Context)
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = natsrpc.Version
var _ = nats_go.Version

type GreeterNATSRPCServer interface {
	HelloToAll(ctx context.Context, req *testdata.HelloRequest)
}

// RegisterGreeterNATSRPCServer register Greeter service
func RegisterGreeterNATSRPCServer(server *natsrpc.Server, s GreeterNATSRPCServer, opts ...natsrpc.ServiceOption) (natsrpc.IService, error) {
	return server.Register("github.com.byebyebruce.natsrpc.example.pb.publish.Greeter", s, opts...)
}

type GreeterNATSRPCClient interface {
	HelloToAll(notify *testdata.HelloRequest, opt ...natsrpc.CallOption) error
}

type _GreeterNATSRPCClient struct {
	c *natsrpc.Client
}

// NewGreeterNATSRPCClient
func NewGreeterNATSRPCClient(enc *nats_go.EncodedConn, opts ...natsrpc.ClientOption) (GreeterNATSRPCClient, error) {
	c, err := natsrpc.NewClient(enc, "github.com.byebyebruce.natsrpc.example.pb.publish.Greeter", opts...)
	if err != nil {
		return nil, err
	}
	ret := &_GreeterNATSRPCClient{
		c: c,
	}
	return ret, nil
}
func (c *_GreeterNATSRPCClient) HelloToAll(notify *testdata.HelloRequest, opt ...natsrpc.CallOption) error {
	return c.c.Publish("HelloToAll", notify, opt...)
}
