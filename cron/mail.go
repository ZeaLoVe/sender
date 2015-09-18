package cron

import (
	"fmt"
	"github.com/open-falcon/sender/g"
	"github.com/open-falcon/sender/model"
	"github.com/open-falcon/sender/proc"
	"github.com/open-falcon/sender/redis"
	"github.com/toolkits/net/httplib"
	"log"
	"net/smtp"
	"strings"
	"time"
)

func Msg(mail *model.Mail) string {
	return fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\r", mail.Tos, mail.Subject, mail.Content)
}

func TosList(mail *model.Mail) []string {
	list := strings.Split(mail.Tos, ",")
	return list
}

//type ACount struct {
//	Server   string
//	User     string
//	Password string
//}

//func SendMail(mail *Mail, acount *ACount) error {
//	auth := smtp.PlainAuth("", acount.User, acount.Password, acount.Server)
//	err := smtp.SendMail(acount.Server+":25", auth, acount.User, mail.TosList(), []byte(mail.Msg()))
//	return err
//}

func ConsumeMail() {
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

func SendMail(mail *model.Mail) {
	defer func() {
		<-MailWorkerChan
	}()
	if g.Config().Maildirect.Enabled {
		auth := smtp.PlainAuth("", g.Config().Maildirect.User, g.Config().Maildirect.Password, g.Config().Maildirect.Server)
		err := smtp.SendMail(g.Config().Maildirect.Server+":25", auth, g.Config().Maildirect.User, TosList(mail), []byte(Msg(mail)))
		if err != nil {
			log.Println(err)
		}
		proc.IncreMailCount()
		return
	}

	url := g.Config().Api.Mail
	r := httplib.Post(url).SetTimeout(5*time.Second, 2*time.Minute)
	r.Param("tos", mail.Tos)
	r.Param("subject", mail.Subject)
	r.Param("content", mail.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	proc.IncreMailCount()

	if g.Config().Debug {
		log.Println("==mail==>>>>", mail)
		log.Println("<<<<==mail==", resp)
	}

}
