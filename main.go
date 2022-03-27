package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic("Redise bağlanılamadı => " + err.Error())
	}
	go listenQueue()
	go addToQueue()

	c := make(chan int)
	<-c
}

func addToQueue() {
	for {
		time.Sleep(time.Second * 1)
		rdb.LPush(ctx, "queue", time.Now().Format("15:04:05"), time.Now().Unix())
	}
}

func listenQueue() {
	for {
		result, err := rdb.BLPop(ctx, 0, "queue").Result()
		if err != nil {
			panic(err)
		}
		if len(result) > 1 {
			for i := 1; i < len(result); i++ {
				println(result[i])
			}
		}
	}
}
