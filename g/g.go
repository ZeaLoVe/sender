package g

import (
	"log"
	"runtime"
)

//v0.0.2 changes:send mail by smtp, set acount in config,add 99u msg ,phone by nexmo api
//v0.0.3 changes:im replace response msg with task_id,use new go-utils im.go
//sdpv0.0.4 go-utils rewrite sendmail, add record phone alarm group, fix bug of TransContent
const (
	VERSION = "sdpv0.0.4"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
