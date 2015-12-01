package cron

import (
	"fmt"
	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/go-utils/smtp"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
	"log"
	"time"
)

var MailSender smtp.SmtpMailSender

func ConsumeMail() {
	if g.Config().Acount.Mail == nil {
		return
	}
	//init mail acount
	var acount smtp.MailAcount
	acount.Password = g.Config().Acount.Mail.Password
	acount.Server = g.Config().Acount.Mail.Server
	acount.User = g.Config().Acount.Mail.User
	MailSender.SetMailAcount(&acount)

	queue := g.Config().Queue.Mail

	for {
		L := redis.PopAllMail(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendMailList(L)
	}
}

func SendMailList(L []*model.Mail) {
	for _, mail := range L {
		MailWorkerChan <- 1
		go SendMail(mail)
	}
}

//sdp function of send mail
//mail acount set by cfg file
func SendMail(mail *model.Mail) {
	defer func() {
		<-MailWorkerChan
	}()

	tmp_mail := smtp.Mail{
		Tos:     mail.Tos,
		Subject: TransContent(mail.Subject),
		Content: TransContent(mail.Content),
	}

	err := MailSender.SendMail(&tmp_mail)
	if err != nil {
		log.Printf("邮件:%s 发送给:%s, 发送失败，错误:%s\n", mail.Subject, mail.Tos, err.Error())
	}
	if err == nil {
		recordMsg := fmt.Sprintf("警报:%s\n已通过 %s 发送给 %s\n",
			mail.Subject,
			"mail",
			mail.Tos)

		RecordAlarm(recordMsg)
	}

	proc.IncreMailCount()

	if g.Config().Debug {
		log.Println("==mail==>>>>", mail)
	}

}
