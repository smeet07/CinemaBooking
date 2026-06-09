package redis
import (
	goredis "github.com/redis/go-redis/v9"
	"context"
	"log")
func NewClient(addr string) *goredis.Client{
	rdb:=goredis.NewClient(&goredis.Options{Addr:addr})
	if err:=rdb.Ping(context.Background()).Err();err!=nil{
		log.Fatalf("Failed to connect to Redis: %v",err)
	}
	log.Printf("Connected to Redis at %s",addr)
	return rdb
}
