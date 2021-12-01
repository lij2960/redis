/************************************************************
 * Author:        jackey
 * Date:        2021/12/1
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"time"
)

var tracer = otel.Tracer("redisexample")

func main()  {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})
	ctx, span := tracer.Start(ctx, "handleRequest")
	defer span.End()
	r, err := rdb.Get(ctx, "sina:ip").Result()
	if err != nil {
		fmt.Println("====", err)
		return
	}
	fmt.Println(r)
	res, err := rdb.Expire(ctx, "sina:ip", 10000000*time.Second).Result()
	if err != nil {
		fmt.Println("====", err)
		return
	}
	fmt.Println("********", res)
}
