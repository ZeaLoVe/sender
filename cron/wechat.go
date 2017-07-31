package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/go-utils/wechat"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
)

func ConsumeWechat() {
	if g.Config().Acount.Wechat == nil {
		return
	}

	corpId := g.Config().Acount.Wechat.CorpId
	agentId := g.Config().Acount.Wechat.AgentId
	secret := g.Config().Acount.Wechat.SecretKey

	wechat.SetConfig(corpId, agentId, secret, "")
	go wechat.GetAccessTokenFromWeixin()
	time.Sleep(5 * time.Second) //等待5S开始，更新token

	queue := g.Config().Queue.Wechat

	for {
		L := redis.PopAllWechat(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendWechatList(L)
	}
}

func SendWechatList(L []*model.WechatSms) {
	for _, wechat := range L {
		WechatWorkerChan <- 1
		go SendWechat(wechat)
	}
}

func SendWechat(sms *model.WechatSms) {
	defer func() {
		<-WechatWorkerChan
	}()

	err := wechat.SendWxMsg(sms.Tos, sms.Content)

	if err != nil {
		log.Printf("微信消息:%s to:%s 发送失败。错误:%s\n", sms.Content, sms.Tos, err.Error())
	}

	if err == nil {
		recordMsg := fmt.Sprintf("警报:%s\n已通过 %s 发送给 %s\n",
			sms.Content,
			"微信",
			sms.Tos)

		RecordAlarm(recordMsg)
	}

	proc.IncreWechatCount()

	if g.Config().Debug {
		log.Println("==wechat==>>>>", sms)
	}
}
