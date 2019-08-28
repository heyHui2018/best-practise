package models

type BindingMQData struct {
	OpenId   string `json:"smarket"` //服务号openid
	MoretvId string `json:"moretvId"`
}

type ProgramMQData struct {
	Sid         string `json:"sid"`
	Title       string `json:"title"`
	ContentType string `json:"contentType"`
}
