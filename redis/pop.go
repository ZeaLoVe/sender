package redis

import (
	"encoding/json"
	"github.com/ZeaLoVe/go-utils/model"
	"github.com/garyburd/redigo/redis"
	"log"
)

func PopAllPhone(queue string) []*model.Phone {
	ret := []*model.Phone{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var phone model.Phone
		err = json.Unmarshal([]byte(reply), &phone)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &phone)
	}

	return ret
}

func PopAllIMSms(queue string) []*model.IMSms {
	ret := []*model.IMSms{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var imsms model.IMSms
		err = json.Unmarshal([]byte(reply), &imsms)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &imsms)
	}

	return ret
}

func PopAllMail(queue string) []*model.Mail {
	ret := []*model.Mail{}

	rc := ConnPool.Get()
	defer rc.Close()

	for {
		reply, err := redis.String(rc.Do("RPOP", queue))
		if err != nil {
			if err != redis.ErrNil {
				log.Println(err)
			}
			break
		}

		if reply == "" || reply == "nil" {
			continue
		}

		var mail model.Mail
		err = json.Unmarshal([]byte(reply), &mail)
		if err != nil {
			log.Println(err, reply)
			continue
		}

		ret = append(ret, &mail)
	}

	return ret
}
