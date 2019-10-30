### 1、.proto 文件生成 .pb.go 文件命令
```
protoc -I . --go_out=plugins=grpc:. route_guide.proto
或
protoc -I proto/ proto/order.proto --go_out=plugins=grpc:order
```
* --go_out用于指定生成源码的保存路径
* -I是-IPATH简写，用于指定查找import文件的路径，可以指定多个
* 最后的order是编译的grpc文件的存储路径
* route_guide.proto为文件名

### 2、.proto 文件内容定义
```
syntax = "proto3";	// 指定语法格式，注意 proto3 不再支持 proto2 的 required 和 optional
package proto;		// 指定生成的 user.pb.go 的包名，防止命名冲突


// service 定义开放调用的服务，即 UserInfoService 微服务
service UserInfoService {
    // rpc 定义服务内的 GetUserInfo 远程调用
    rpc GetUserInfo (UserRequest) returns (UserResponse) {
    }
}

// message 对应生成代码的 struct。
// 如下面的UserRequest,即为上面方法GetUserInfo的请求,UserResponse即为上面方法GetUserInfo的返回

// 定义客户端请求的数据格式
message UserRequest {
	// [修饰符] 类型 字段名 = 标识符;
	string name = 1;
}

// 定义服务端响应的数据格式
message UserResponse {
    int32 id = 1;
    string name = 2;
    int32 age = 3;
    repeated string title = 4;	// repeated 修饰符表示字段是可变数组，即 slice 类型
    (required表示该字段值是必须设置的;optional表示该字段的值可以存在,也可以为空)
}
```
更多字段类型可参考：https://colobu.com/2015/01/07/Protobuf-language-guide/#%E6%A0%87%E9%87%8F%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B

### 3、.pb.go 文件内容说明
```
package proto

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// 请求结构
type UserRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

// 为字段自动生成的 Getter
func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 响应结构
type UserResponse struct {
	Id    int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age   int32    `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Title []string `protobuf:"bytes,4,rep,name=title" json:"title,omitempty"`
}

// 客户端需实现的接口
type UserInfoServiceClient interface {
	GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

// 服务端需实现的接口
type UserInfoServiceServer interface {
	GetUserInfo(context.Context, *UserRequest) (*UserResponse, error)
}

// 将微服务注册到 grpc 
func RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer) {
	s.RegisterService(&_UserInfoService_serviceDesc, srv)
}
// 处理请求
func _UserInfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {...}
```

### 4、message更新规则
* 不可更改已存在域中的标识号
* 所有新增的域必须是optional或repeated
* 非required域可以被删除,但这些被删除域的标识号不可再被使用
* 非required域可以被转化,转化时可能发生扩展或截断,此时标识号和名称均不变
* sint32和sint64是互相兼容的
* fixed32兼容sfixed32,fixed64兼容sfixed64
* optional兼容repeated,发送端发送repeated域,用户使用optional域读取,将会读取repeated域的最后一个元素

### 5、序列化原理
Varint：是一种紧凑的数字表示方法,用一个或多个字节来完成,值越小的数字用的字节数越少.
其中每个byte的最高位有特殊的含义,若为1,表示后续的字节也是该数字的一部分,若为0,则结束,其余7为用来表示数字.
(未完)
