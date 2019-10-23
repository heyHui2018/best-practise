package rpc

import (
	"github.com/heyHui2018/best-practise/pb"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"io"
	"time"
)

func (this *Server) Communicate(stream pb.User_CommunicateServer) error {
	t := new(log.TLog)
	t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			t.Infof("Communicate,EOF,read done")
			return nil
		}
		if err != nil {
			t.Warnf("Communicate,stream.Recv error,err = %v", err)
			return err
		}
		t.Infof("Communicate,in = %v", in)
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}

/*
client:
package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/heyHui2018/best-practise/pb"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"google.golang.org/grpc"
	"io"
	"time"
)

func Reply(client pb.UserClient) {
	t := new(log.TLog)
	t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	m := &pb.UserInfoRequest{Greet: "hello"}
	out, err := proto.Marshal(m)
	if err != nil {
		t.Warnf("Reply,proto.Marshal error,err = %v", err)
		return
	}
	notes := []*pb.Message{
		{Type: "UserInfoRequest", Data: out},
	}
	stream, err := client.Communicate(context.Background())
	if err != nil {
		t.Warnf("Reply,client.Communicate error,err = %v", err)
		return
	}
	waitc := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Second)
			in, err := stream.Recv()
			if err == io.EOF {
				t.Infof("Reply,EOF,read done")
				close(waitc)
				return
			}
			if err != nil {
				t.Warnf("Reply,stream.Recv error,err = %v", err)
				return
			}
			t.Infof("Reply,in = %+v", in)
			if in.Type == "UserInfoRequest" {
				mess := &pb.UserInfoRequest{}
				if err := proto.Unmarshal(in.Data, mess); err != nil {
					t.Warnf("Reply, proto.Unmarshal error,err = %v", err)
				}
				t.Infof("Reply,mess = %+v", mess)
			}
			stream.Send(notes[0])
		}
	}()
	t.Infof("Reply,notes = %v", notes)
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			t.Warnf("Reply,stream.Send error,err = %v", err)
		}
	}
	// stream.CloseSend()
	<-waitc
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:8667", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	k := pb.NewUserClient(conn)
	Reply(k)
}
// 运行此client可能会报关于conn类型的错,可通过删除此包的vendor中grpc的包解决
*/
