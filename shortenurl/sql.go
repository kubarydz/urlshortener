package shortenurl

import (
	"github.com/redis/go-redis/v9"
)

func Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})
}
