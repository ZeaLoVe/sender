package proc

import (
	"sync/atomic"
)

var imsmsCount, mailCount, phoneCount uint32

func GetIMSmsCount() uint32 {
	return atomic.LoadUint32(&imsmsCount)
}

func GetMailCount() uint32 {
	return atomic.LoadUint32(&mailCount)
}

func GetPhoneCount() uint32 {
	return atomic.LoadUint32(&phoneCount)
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
