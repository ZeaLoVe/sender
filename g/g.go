package g

import (
	"log"
	"runtime"
)

//v0.0.2 changes
//send mail by smtp, set acount in config
//add 99u msg ,phone by nexmo api
const (
	VERSION = "sdpv0.0.2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
