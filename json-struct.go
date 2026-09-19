package main

import (
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func SaveUser(client *redis.Client, user User) {
	//user marshal when struct to json and use unmarshal when json to struct
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, "user:101", jsonBytes, time.Hour)
}

func GetUser(client *redis.Client, id string) *User {
	jsonString, err := client.Get(ctx, "user:"+id).Result()
	if err == redis.Nil {
		return nil
	}

	var user User
	err = json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		panic(err)
	}
	return &user
}
