package cron

import (
	"github.com/ZeaLoVe/go-utils/im"
	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
	"log"
	"strings"
	"time"
)

var IMSender im.IM99U

func ConsumeIMSms() {
	ac := im.Acount{
		Uri:      g.Config().Acount.IM.Uri,
		Password: g.Config().Acount.IM.Password,
	}
	IMSender.SetAcount(&ac)

	queue := g.Config().Queue.IMSms
	for {
		L := redis.PopAllIMSms(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendSmsList(L)
	}
}

func SendSmsList(L []*model.IMSms) {
	for _, imsms := range L {
		IMSmsWorkerChan <- 1
		go SendSms(imsms)
	}
}

func SendSms(imsms *model.IMSms) {
	defer func() {
		<-IMSmsWorkerChan
	}()

	//	fmt.Println(sms.Tos)
	IMSender.SendMsg(strings.Split(imsms.Tos, ","), imsms.Content)

	proc.IncreIMSmsCount()

	if g.Config().Debug {
		log.Println("==im-sms==>>>>", imsms)
	}
}
