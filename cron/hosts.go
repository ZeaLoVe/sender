package cron

import (
	"sync"
	"net/http"
	"github.com/ZeaLoVe/sender/g"
)

type ResponseHost struct {
	Ip       string `json:"ip,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}

//endpoint-ip  kv
type SafeHostMap struct {
	sync.Mutex
	M map[string]string
}

var HostMap = &SafeHostMap{M: make(map[string]string)}

func (this *SafeHostMap) GetIP(endpoint string) string {
	this.Lock()
	defer this.Lock()
	return this.M[endpoint]
}

func (this *SafeHostMap) UpdateMap(hosts []ResponseHost) {
	this.Lock()
	defer this.Lock()
	for _, host := range hosts {
		this.M[host.Endpoint] = host.Ip
	}
}

func GetHostMap() host []ResponseHost{
	http.Get()
	
}

func UpdateHostMap(){
	
}