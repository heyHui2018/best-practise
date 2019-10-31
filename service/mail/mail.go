package mail

import (
	"encoding/json"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/heyHui2018/log"

	"github.com/heyHui2018/best-practise/base"
	dataSource2 "github.com/heyHui2018/best-practise/model/dataSource"
	"github.com/heyHui2018/best-practise/service/dataSource"
)

var MailHostMap = map[string]string{
	"sina":    "smtp.sina.com:25",      // 新浪
	"qq":      "smtp.qq.com:25",        // qq
	"tencent": "smtp.exmail.qq.com:25", // 腾讯企业
	"126":     "smtp.126.com:25",       // 126
	"163":     "smtp.163.com:25",       // 163
	"gmail":   "smtp.gmail.com:25",     // gmail
}

func SendMail(t *log.TLog) {
	// 从数据库查询当前时刻需发送邮件的用户
	rr := new(dataSource2.RegisterRecord)
	rr.Hour = strconv.Itoa(time.Now().Hour())
	rrList, err := rr.FindByHour()
	if err != nil {
		t.Warnf("SendMail FindByHour error,err = %v", err)
		return
	}
	to := ""
	for k := range rrList {
		if k == len(rrList)-1 {
			to += rrList[k].Email
		} else {
			to += rrList[k].Email + ";"
		}
	}
	mail := base.GetConfig().Mail
	// todo 自动识别邮箱所属host/contentType可自选/body优化
	auth := smtp.PlainAuth("", mail.Username, mail.Password, strings.Split(MailHostMap["tencent"], ":")[0])
	var contentType string
	if mail.MailType == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	subject := "Weather Today"
	data := dataSource.GetWeather(t, "ShangHai", "ShangHai", "China")
	body, err := json.Marshal(data)
	if err != nil {
		t.Warnf("SendMail Marshal error,err = %v", err)
		return
	}
	msg := []byte("To: " + to + "\r\nFrom: " + mail.Nickname + "<" + mail.Username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n")
	msg = append(msg, body...)
	err = smtp.SendMail(MailHostMap["tencent"], auth, mail.Username, strings.Split(to, ";"), msg)
	if err != nil {
		t.Warnf("SendMail error,err = %v", err)
		return
	}
	t.Infof("SendMail 完成")
}
