package cron

import (
	"github.com/ZeaLoVe/sender/g"
)

var (
	IMSmsWorkerChan chan int
	MailWorkerChan  chan int
	PhoneWorkerChan chan int
)

func InitWorker() {
	workerConfig := g.Config().Worker
	IMSmsWorkerChan = make(chan int, workerConfig.IMSms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
	PhoneWorkerChan = make(chan int, workerConfig.Phone)
}
