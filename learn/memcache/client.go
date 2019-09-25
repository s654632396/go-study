package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

	mc := memcache.New("127.0.0.1:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	it, err := mc.Get("foo")
	if err != nil {
		println("err:", err)
	} else {
		println("item:", it)
	}
}
