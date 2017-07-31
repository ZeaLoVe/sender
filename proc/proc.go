package proc

import (
	"sync/atomic"
)

var imsmsCount, mailCount, phoneCount, wechatCount uint32

func GetIMSmsCount() uint32 {
	return atomic.LoadUint32(&imsmsCount)
}

func GetMailCount() uint32 {
	return atomic.LoadUint32(&mailCount)
}

func GetPhoneCount() uint32 {
	return atomic.LoadUint32(&phoneCount)
}

func GetWechatCount() uint32 {
	return atomic.LoadUint32(&wechatCount)
}

func IncreIMSmsCount() {
	atomic.AddUint32(&imsmsCount, 1)
}

func IncreMailCount() {
	atomic.AddUint32(&mailCount, 1)
}

func IncrePhoneCount() {
	atomic.AddUint32(&phoneCount, 1)
}

func IncreWechatCount() {
	atomic.AddUint32(&wechatCount, 1)
}
