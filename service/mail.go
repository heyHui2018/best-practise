package service

import (
	"encoding/json"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models"
	"github.com/ngaut/log"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

var MailHostMap = map[string]string{
	"sina":    "smtp.sina.com:25",      // 新浪
	"qq":      "smtp.qq.com:25",        // qq
	"tencent": "smtp.exmail.qq.com:25", // 腾讯企业
	"126":     "smtp.126.com:25",       // 126
	"163":     "smtp.163.com:25",       // 163
	"gmail":   "smtp.gmail.com:25",     // gmail
}

func SendMail(traceId string) {
	// 从数据库查询当前时刻需发送邮件的用户
	rr := new(models.RegisterRecord)
	rr.Hour = strconv.Itoa(time.Now().Hour())
	rrList, err := rr.FindByHour()
	if err != nil {
		log.Warnf("SendMail FindByHour error,traceId = %v,err = %v", traceId, err)
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
	data := GetWeather(traceId, "ShangHai", "ShangHai", "China")
	body, err := json.Marshal(data)
	if err != nil {
		log.Warnf("SendMail Marshal error,traceId = %v,err = %v", traceId, err)
		return
	}
	msg := []byte("To: " + to + "\r\nFrom: " + mail.Nickname + "<" + mail.Username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n")
	msg = append(msg, body...)
	log.Infof("body = %v", string(body))
	err = smtp.SendMail(MailHostMap["tencent"], auth, mail.Username, strings.Split(to, ";"), msg)
	if err != nil {
		log.Warnf("SendMail error,err = %v", err)
		return
	}
	log.Infof("SendMail 完成,traceId = %v", traceId)
}
