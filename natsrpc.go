package natsrpc

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/byebyebruce/natsrpc/encode/pb"
)

var (
	ErrHeaderFormat     = errors.New("natsrpc: header format error")
	ErrDuplicateService = errors.New("natsrpc: duplicate service")
	ErrNoMethod         = errors.New("natsrpc: no method")
	ErrNoMeta           = errors.New("natsrpc: no meta data")
	ErrEmptyReply       = errors.New("natsrpc: reply is empty")

	// ErrReplyLater
	// It's not an error, when you want to reply message later, then return this.
	ErrReplyLater = errors.New("natsrpc: reply later")
)

const (
	defaultSubQueue = "_ns_q" // default queue group

	headerMethod = "_ns_method" // method
	headerUser   = "_ns_user"   // user header
	headerError  = "_ns_error"  // reply error
	pubSuffix    = "_pub"       // publish subject suffix
)

// ServiceRegistrar 注册服务
type ServiceRegistrar interface {
	// Register 注册
	Register(sd ServiceDesc, svc any, opt ...ServiceOption) (IService, error)
}

// IClient 客户端接口
type IClient interface {
	// Publish 发布
	Publish(method string, req interface{}, opt ...CallOption) error

	// Request 请求
	Request(ctx context.Context, method string, req interface{}, rep interface{}, opt ...CallOption) error
}

// IService 服务
type IService interface {
	// Name 名字
	Name() string

	// Close 关闭
	Close() bool
}

// Encoder 编码器
type Encoder interface {
	// Encode 编码
	Encode(v interface{}) ([]byte, error)

	// Decode 解码
	Decode(data []byte, vPtr interface{}) error
}

// DefaultServerOptions 默认server选项
var DefaultServerOptions = ServerOptions{
	errorHandler: func(i interface{}) {
		fmt.Fprintf(os.Stderr, "error:%v\n", i)
	},
	recoverHandler: func(i interface{}) {
		fmt.Fprintf(os.Stderr, "server panic:%v\n", i)
	},
	encoder: pb.Encoder{},
}

// DefaultServiceOptions 默认service选项
var DefaultServiceOptions = ServiceOptions{
	timeout:        time.Duration(5) * time.Second,
	multiGoroutine: false,
	id:             "",
}

// DefaultClientOptions 默认client选项
var DefaultClientOptions = ClientOptions{
	namespace: "",
	id:        "",
	encoder:   pb.Encoder{},
}
