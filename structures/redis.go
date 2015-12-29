package structures

import(
	"github.com/garyburd/redigo/redis"
	"log"
)

type RedisAccessInfo struct{
	Pool *redis.Pool
}


func NewRedis() *RedisAccessInfo{

	mypool := redis.Pool{
		MaxIdle : 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error){
			c, err := redis.Dial("tcp", ":6379")
			if err != nil{
				panic(err.Error())
			}else{
				log.Println("Redis Connection Initialized")
			}
			return c,err
		},
	}
	log.Println("Redis Pool Initalized")
	return &RedisAccessInfo{
		Pool : &mypool,
	}
}
