package cron

import (
	"github.com/ZeaLoVe/go-utils/model"
	"github.com/ZeaLoVe/go-utils/nexmo"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/proc"
	"github.com/ZeaLoVe/sender/redis"
	"log"
	"strings"
	"time"
)

var PhoneSender nexmo.Nexmo

func ConsumePhone() {

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

	for _, to := range strings.Split(phone.Tos, ",") {
		PhoneSender.SetToChinaZoneCode(to)
		PhoneSender.SetVoiceMsg(phone.Content)
		err := PhoneSender.Call()
		if err != nil {
			log.Println(err.Error())
		}

		proc.IncrePhoneCount()
		if g.Config().Debug {
			log.Println("==phone==>>>>", phone)
		}
	}

}
