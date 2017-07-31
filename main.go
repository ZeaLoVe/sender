package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZeaLoVe/sender/cron"
	"github.com/ZeaLoVe/sender/g"
	"github.com/ZeaLoVe/sender/http"
	"github.com/ZeaLoVe/sender/redis"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	cron.InitWorker()
	redis.InitConnPool()

	go http.Start()
	go cron.ConsumeIMSms()
	go cron.ConsumeMail()
	go cron.ConsumePhone()
	go cron.ConsumeWechat()
	go cron.UpdateHostMap()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		redis.ConnPool.Close()
		os.Exit(0)
	}()

	select {}
}
