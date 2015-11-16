package http

import (
	"fmt"
	"github.com/ZeaLoVe/sender/proc"
	"net/http"
)

func configProcRoutes() {

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("im_sms:%v, mail:%v, phone:%v", proc.GetIMSmsCount(), proc.GetMailCount(), proc.GetPhoneCount())))
	})

}
