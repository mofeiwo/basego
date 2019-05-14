package exdt

import (
	"testing"
	"fmt"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	redisClient = new(RedisClient).getInstance(map[string]string{
		"addr":     "127.0.0.1:6379",
		"password": "",
		"DB":       "0",
	})
}

func TestGetStringValue(t *testing.T) {
	redisClient.Set("redisname", "test", 0)
	if _, err := redisClient.Get("redisname").Result(); err != nil {
		t.Error("redis client error:", err)
	}

	fmt.Println(redisClient.Del("redisname").Result())
}

func BenchmarkMySQLClient_Instance(b *testing.B) {
	for i := 0; i < 10; i++ {
		redisClient.Set("redisname", "test", 0)
	}
	fmt.Println("test success!")
}

func TestMySQLClient_Instance(t *testing.T) {

}
