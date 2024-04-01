package main

import (
	"fmt"
	"time"

	"github.com/erni27/imcache"
)

func LogEvictedEntry(key string, value interface{}, reason imcache.EvictionReason) {
	fmt.Printf("Evicted entry: %s=%v (%s)", key, value, reason)
	//log.Printf("Evicted entry: %s=%v (%s)", key, value, reason)
}

func main() {
	var c *imcache.Cache[string, interface{}]
	c = imcache.New[string, interface{}](
	//imcache.WithDefaultExpirationOption[string, interface{}](time.Second),
	//imcache.WithNoExpiration()
	//imcache.WithEvictionCallbackOption[string, interface{}](LogEvictedEntry),
	)
	//c.Set("foo", "bar", imcache.WithDefaultExpiration())
	c.Set("foo", "bar", imcache.WithDefaultExpiration())
	c.Set("1", "2", imcache.WithDefaultExpiration())
	c.Set("3", "ar", imcache.WithDefaultExpiration())

	v, ok := c.Get("foo")
	if ok {
		fmt.Println("expected entry to be expired", v)
	} else {
		fmt.Println(v)
	}

	time.Sleep(2 * time.Second)

	v, ok = c.Get("foo")
	if ok {
		fmt.Println("expected entry to be expired", v)
	} else {
		fmt.Println(v)
	}

	g := c.GetAll()
	delete(g, "3")

	fmt.Println(g)
	fmt.Println(c.GetAll())
}
