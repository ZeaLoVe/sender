package cron

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ZeaLoVe/go-utils/im"
	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
)

var IMSender im.IM99U

func RecordAlarm(recordMsg string) {
	tos := g.Config().Acount.IM.Group
	if tos == "" {
		return
	}
	err := IMSender.SendMsg(strings.Split(tos, ","), recordMsg)
	if err != nil {
		log.Printf("record fail:%s with err:%s", recordMsg, err.Error())
	}
}

func RecordPhoneAlarm(recordMsg string) {
	tos := g.Config().Acount.IM.PhoneGroup
	if tos == "" {
		return
	}
	err := IMSender.SendMsg(strings.Split(tos, ","), recordMsg)
	if err != nil {
		log.Printf("record phone alarm fail:%s with err:%s", recordMsg, err.Error())
	}
}

func ConsumeIMSms() {
	if g.Config().Acount.IM == nil {
		return
	}
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
	err := IMSender.SendMsg(strings.Split(imsms.Tos, ","), TransContent(imsms.Content))

	if err != nil {
		log.Printf("IM消息:%s to:%s 发送失败。错误:%s\n", imsms.Content, imsms.Tos, err.Error())
	}

	if err == nil {
		recordMsg := fmt.Sprintf("警报:%s\n已通过 %s 发送给 %s\n",
			imsms.Content,
			"99u",
			imsms.Tos)

		RecordAlarm(recordMsg)
	}

	proc.IncreIMSmsCount()

	if g.Config().Debug {
		log.Println("==im-sms==>>>>", imsms)
	}
}
