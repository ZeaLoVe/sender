package model

import (
	"fmt"
)

type Sms struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
}

type Mail struct {
	Tos     string `json:"tos"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (mail *Mail) Msg() string {
	return fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\r", mail.Tos, mail.Subject, mail.Content)
}

func (mail *Mail) TosList() []string {
	list := strings.Split(mail.Tos, ",")
	return list
}

type ACount struct {
	Server   string
	User     string
	Password string
}

func SendMail(mail *Mail, acount *ACount) error {
	auth := smtp.PlainAuth("", acount.User, acount.Password, acount.Server)
	err := smtp.SendMail(acount.Server+":25", auth, acount.User, mail.TosList(), []byte(mail.Msg()))
	return err
}

func (this *Sms) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Content:%s>",
		this.Tos,
		this.Content,
	)
}

func (this *Mail) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Subject:%s, Content:%s>",
		this.Tos,
		this.Subject,
		this.Content,
	)
}
