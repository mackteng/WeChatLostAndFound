package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
)

type Redis struct {
	Pool *redis.Pool
}

func NewRedis() *Redis {

	mypool := redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			} else {
				log.Println("Redis Connection Initialized")
			}
			return c, err
		},
	}
	log.Println("Redis Pool Initalized")
	return &Redis{
		Pool: &mypool,
	}
}

func (Redis *Redis) AddMessageToQueue(OpenID string, Channel int, Payload string) error {

	key := "User:" + OpenID + ":Channel:" + strconv.Itoa(Channel)
	conn := Redis.Pool.Get()

	defer conn.Close()

	_, err := conn.Do("LPUSH", key, Payload)
	_, err = conn.Do("LTRIM", key, 0, 4)
	log.Println("Added Message to Queue : " + OpenID)
	return err
}

func (Redis *Redis) GetMessagesFromQueue(OpenID string, Channel int) ([]string, error) {

	key := "User:" + OpenID + ":Channel:" + strconv.Itoa(Channel)
	conn := Redis.Pool.Get()
	defer conn.Close()

	res, err := conn.Do("LRANGE", key, 0, -1)

	if err == nil {
		if strs, err := redis.Strings(res, err); err == nil {
			return strs, nil
		}
	} else {

		return nil, err
	}

	log.Println("Flushing Channel ", Channel, " of ", OpenID)
	return nil, err
}

func (Redis *Redis) insertMsgID(MsgID string) error {

	conn := Redis.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", MsgID, true)
	_, err = conn.Do("EXPIRE", MsgID, 30)
	if err != nil {
		return err
	}

	log.Println("Added MsgID " + MsgID)
	return nil
}

func (Redis *Redis) IsDuplicateMsgID(MsgID string) (bool, error) {

	conn := Redis.Pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", MsgID))

	if exists {
		log.Println("Duplicate MessageID! Ignoring!")
		return true, err
	}
	return false, Redis.insertMsgID(MsgID)

}
