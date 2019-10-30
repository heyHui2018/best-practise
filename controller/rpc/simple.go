package rpc

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/pb"
)

type Server struct{}

func (s *Server) GetTimestamp(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	t := new(log.TLog)
	t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	start := time.Now()
	data := make(map[string]interface{})
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	ymd := strings.Split(timeStr, "-")
	data["year"] = ymd[0]
	data["month"] = ymd[1]
	data["day"] = ymd[2]
	r := new(pb.GetReply)
	dataStr, err := json.Marshal(data)
	if err != nil {
		t.Warnf("GetTimestamp error,err = %v", err)
		return nil, err
	}
	r.Status = base.Success
	r.Message = base.CodeText[base.Success]
	r.Timestamp = time.Now().Unix()
	r.Data = string(dataStr)
	t.Infof("GetTimestamp 完成,res = %v,耗时 = %v", r, time.Since(start))
	return r, nil
}

// client demo : https://github.com/heyHui2018/demo/tree/master/rpc/grpc/getTimestampDemo/client
