package cache

import (
	"github.com/go-redis/redis"
)

var Cache *redis.Client

func InitCache() (err error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = db.Ping().Result()
	if err != nil {
		return err
	}
	Cache = db
	return nil
}
