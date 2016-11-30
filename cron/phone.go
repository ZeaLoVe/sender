package cron

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/go-utils/nexmo"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
)

var PhoneSender nexmo.Nexmo

func ConsumePhone() {

	if g.Config().Acount.Phone == nil {
		return
	}

	PhoneSender.SetKeyAndSecret(g.Config().Acount.Phone.Key, g.Config().Acount.Phone.Serect)
	//	nexmo.Status_url = "http://61.147.187.235:8011/hello"

	queue := g.Config().Queue.Phone
	for {
		L := redis.PopAllPhone(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendPhoneList(L)
	}
}

func SendPhoneList(L []*model.Phone) {
	for _, phone := range L {
		PhoneWorkerChan <- 1
		go SendPhone(phone)
	}
}

func SendPhone(phone *model.Phone) {
	defer func() {
		<-PhoneWorkerChan
	}()

	//	fmt.Println(sms.Tos)
	PhoneSender.SetRepeat("2")
	PhoneSender.SetLanguage("zh-cn")
	if g.Config().Acount.Phone.Callback != "" {
		PhoneSender.SetCallBack(g.Config().Acount.Phone.Callback)
	}

	for _, to := range strings.Split(phone.Tos, ",") {
		if len(to) == 0 {
			continue
		}
		PhoneSender.SetToChinaZoneCode(to)
		PhoneSender.SetVoiceMsg(TransContent(phone.Content))
		err, resp := PhoneSender.Call()
		if err != nil {
			log.Printf("报警电话:%s to:%s 调用API失败,错误:%s\n", phone.Content, to, err.Error())
		}
		if err == nil {
			recordMsg := fmt.Sprintf("警报:%s\n已通过 %s 发送给 %s 呼叫编码%s\n",
				phone.Content,
				"phone",
				to,
				resp.Call_id)

			RecordAlarm(recordMsg)
			RecordPhoneAlarm(recordMsg)
		}

		proc.IncrePhoneCount()

	}
	if g.Config().Debug {
		log.Println("==phone==>>>>", phone)
	}
}
