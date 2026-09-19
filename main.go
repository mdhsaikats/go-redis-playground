package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// so we need init a context cause it will tell the redis when we need to froze the connection
var ctx = context.Background()

func main() {
	//then we need to init a client that will tell the redis address,pass,db
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "29112003",
		DB:       0, // total have 16 database
	})

	//if we want something we have to add the .Result() with that function
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to the redis: %v", err)
	}

	fmt.Println("Connected to redis!", pong)

	//set function used to set key and value to that key and add TTL for that key and the value
	err = client.Set(ctx, "username", "saikat", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	//get function to get the saved value from the redis database
	val, err := client.Get(ctx, "username").Result()

	if err == redis.Nil {
		fmt.Println("Key dosent exists in redis anymore!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Found username", val)
	}

	SaveUser(client, User{ID: 10, Email: "Saikat"})
	fmt.Println(GetUser(client, "101"))
}
