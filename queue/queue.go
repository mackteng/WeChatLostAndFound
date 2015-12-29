package queue

import(
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"log"
	
)


func AddMessageToQueue(OpenID string, Channel int, Payload string, redisaccess *structures.RedisAccessInfo) error {

	key := "User:" + OpenID + ":Channel:" + strconv.Itoa(Channel)
	conn := redisaccess.Pool.Get()
	
	defer conn.Close()

	res, err := conn.Do("LPUSH", key, Payload)
	res, err = conn.Do("LTRIM", key, 0, 4)
	log.Println(res)
	return err
}



func Flush(OpenID string, Channel int, redisaccess *structures.RedisAccessInfo) ([]string, error) {

	log.Println("Flushing Channel ", Channel, " of ", OpenID)

	key := "User:" + OpenID + ":Channel:" + strconv.Itoa(Channel)
        conn := redisaccess.Pool.Get()

        defer conn.Close()

	res,err := conn.Do("LRANGE", key, 0, -1)

	if err == nil{
		if strs, err := redis.Strings(res,err); err==nil {
			log.Println(strs)
			return strs, nil
		}
	} else {

		log.Println(err)
	}

	return nil, err
}
