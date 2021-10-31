// This code was autogenerated from natsrpc.gencode, do not edit.
package pb

import (
	"context"

	"github.com/byebyebruce/natsrpc"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

var _ proto.Message

// Greeter
type Greeter interface {
	// AreYouOK1
	AreYouOK1(ctx context.Context, req *HelloRequest) (*HelloReply, error)
	// AreYouOK2Async
	AreYouOK2Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error))
	// AreYouOK3Async
	AreYouOK3Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error))
	// PublishAreYouOK4
	PublishAreYouOK4(ctx context.Context, notify *HelloRequest) error
}

// RegisterGreeter
func RegisterGreeter(server *natsrpc.Server, s Greeter, opts ...natsrpc.Option) (natsrpc.Service, error) {
	return server.Register(".;pb.Greeter", s, opts...)
}

// GreeterClient
type GreeterClient struct {
	c *natsrpc.Client
}

// NewGreeterClient
func NewGreeterClient(enc *nats.EncodedConn, opts ...natsrpc.Option) (*GreeterClient, error) {
	c, err := natsrpc.NewClient(enc, ".;pb.Greeter", opts...)
	if err != nil {
		return nil, err
	}
	ret := &GreeterClient{
		c: c,
	}
	return ret, nil
}

// ID 根据ID获得client
func (c *GreeterClient) ID(id interface{}) *GreeterClient {
	return &GreeterClient{
		c: c.c.ID(id),
	}
}

// AreYouOK1
func (c *GreeterClient) AreYouOK1(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	rep := &HelloReply{}
	err := c.c.Request(ctx, "AreYouOK1", req, rep)
	return rep, err
}

// AreYouOK2Async
func (c *GreeterClient) AreYouOK2Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error)) {
	rep := &HelloReply{}
	f := func(_ proto.Message, err error) {
		cb(rep, err)
	}
	c.c.AsyncRequest("AreYouOK2", req, rep, f)
}

// AreYouOK3Async
func (c *GreeterClient) AreYouOK3Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error)) {
	rep := &HelloReply{}
	f := func(_ proto.Message, err error) {
		cb(rep, err)
	}
	c.c.AsyncRequest("AreYouOK3", req, rep, f)
}

// PublishAreYouOK4
func (c *GreeterClient) PublishAreYouOK4(ctx context.Context, notify *HelloRequest) error {
	return c.c.Publish("AreYouOK4", notify)
}

// AsyncGreeter
type AsyncGreeter interface {
	// AreYouOK1Async
	AreYouOK1Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error))
	// AreYouOK2Async
	AreYouOK2Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error))
	// AreYouOK3Async
	AreYouOK3Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error))
}

// RegisterAsyncGreeter
func RegisterAsyncGreeter(server *natsrpc.Server, s AsyncGreeter, opts ...natsrpc.Option) (natsrpc.Service, error) {
	return server.Register(".;pb.AsyncGreeter", s, opts...)
}

// AsyncGreeterClient
type AsyncGreeterClient struct {
	c *natsrpc.Client
}

// NewAsyncGreeterClient
func NewAsyncGreeterClient(enc *nats.EncodedConn, opts ...natsrpc.Option) (*AsyncGreeterClient, error) {
	c, err := natsrpc.NewClient(enc, ".;pb.AsyncGreeter", opts...)
	if err != nil {
		return nil, err
	}
	ret := &AsyncGreeterClient{
		c: c,
	}
	return ret, nil
}

// ID 根据ID获得client
func (c *AsyncGreeterClient) ID(id interface{}) *AsyncGreeterClient {
	return &AsyncGreeterClient{
		c: c.c.ID(id),
	}
}

// AreYouOK1Async
func (c *AsyncGreeterClient) AreYouOK1Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error)) {
	rep := &HelloReply{}
	f := func(_ proto.Message, err error) {
		cb(rep, err)
	}
	c.c.AsyncRequest("AreYouOK1", req, rep, f)
}

// AreYouOK2Async
func (c *AsyncGreeterClient) AreYouOK2Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error)) {
	rep := &HelloReply{}
	f := func(_ proto.Message, err error) {
		cb(rep, err)
	}
	c.c.AsyncRequest("AreYouOK2", req, rep, f)
}

// AreYouOK3Async
func (c *AsyncGreeterClient) AreYouOK3Async(ctx context.Context, req *HelloRequest, cb func(*HelloReply, error)) {
	rep := &HelloReply{}
	f := func(_ proto.Message, err error) {
		cb(rep, err)
	}
	c.c.AsyncRequest("AreYouOK3", req, rep, f)
}
