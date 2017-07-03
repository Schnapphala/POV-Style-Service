package Memcache

import (
	gomemcache "github.com/bradfitz/gomemcache/memcache"
	"time"
)


type Connection struct {
	Ip string
	Port uint16
}

type Memcache struct {
	Connection Connection
	Client *gomemcache.Client
}

func NewMemcache (ip string, port uint16) Memcache {
	var connection = Connection{Ip : ip, Port:port }
	return Memcache{Connection : connection};
}

func (memcache Memcache) connect () {
	memcache.Client = gomemcache.New(memcache.Connection.Ip + ":" + string(memcache.Connection.Port))
}


func (memcache Memcache) get(key string) (string, bool) {

	if memcache.Client != nil {
		item, error := memcache.Client.Get(key)

		if error != nil {
			return string(item.Value), true
		}
	}

	return "", false;
}

func (memcache Memcache) set(key, value string) bool {

	if memcache.Client != nil {
		item := gomemcache.Item{
			Value:[]byte(value),
			Key:key,
			Expiration:time.Now().Unix()
		}


		memcache.Client.Set()
	}
}