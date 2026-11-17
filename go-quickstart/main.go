package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := client.Set(ctx, "bike:1", "Process 134", 0).Err(); err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "bike:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
