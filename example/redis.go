package main

import (
	"fmt"

	"github.com/quangdangfit/gocommon/redis"
)

func main() {
	var conf = redis.Config{
		Address:  "localhost:6379",
		Password: "password",
		Database: 1,
	}
	var r = redis.New(conf)

	if err := r.Set("key", "value"); err != nil {
		fmt.Println(err)
	}

	var value string
	if err := r.Get("key", &value); err != nil {
		fmt.Println(err)
	}
	fmt.Println("value: ", value)
}
