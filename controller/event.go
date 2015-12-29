package controller

import (
	"log"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/wechat"
	"bitbucket.org/mack_teng/WeChatLostAndFound/queue"
	"errors"
	"time"
	"strconv"
)

var entryhandlers = map[string]handler{
	"RegisterTag":RegisterTag,
	"FindTag" : FindTag,
	
	"1": changeChannel,	
	"2": changeChannel,	
	"3": changeChannel,	
	"4": changeChannel,	
	"5": changeChannel,

	"6": changeChannel,
	"7": changeChannel,
	"8": changeChannel,
	"9": changeChannel,
	"10": changeChannel,
	
}


func EventHandler(q *structures.Message, config *structures.GlobalConfiguration) error {

	log.Println("EventMessageHandlerCalled", q)
	err := entryhandlers[q.EventKey](q,config)
	
	if err!=nil {
		log.Println(err)
	}

	return nil
}


func RegisterTag(q *structures.Message, config *structures.GlobalConfiguration) error{

	log.Println("Controller RegisterTag Called", q)

	test := structures.ItemInfo{
		q.ScanCodeInfo.ScanResult,
		"foo",
		"foo",
	}

	dbconfig := config.DatabaseConfig
	OpenID := q.FromUserName
	var err error

	if !database.UserExists(dbconfig, OpenID) {
                database.AddUser(dbconfig, OpenID)
        }

        if database.ItemExists(dbconfig, test.TagID) {
		wechat.SendItemAlreadyRegistered(OpenID, config)
                return errors.New("Item Already Registered")
        }

        next_channel := database.NextOwnerChannel(dbconfig, OpenID)

        if next_channel < 0 {
                return errors.New(structures.REGISTER_LIMIT)
        }


	err = database.RegisterTag(config.DatabaseConfig, q.FromUserName, next_channel, &test)

	if err == nil{
		err = database.ChangeChannel(dbconfig, OpenID, next_channel)
		if err == nil{
			wechat.SendChannelChangeConfirmation(OpenID, next_channel,config)
		}
	}	

	return err
}


func FindTag(q *structures.Message, config *structures.GlobalConfiguration) error{

	log.Println("FindTag Called", q)
	dbconfig := config.DatabaseConfig
	FinderOpenID := q.FromUserName
	TagID := q.ScanCodeInfo.ScanResult

	if !database.UserExists(dbconfig, FinderOpenID){
                database.AddUser(dbconfig, FinderOpenID)
        }

        if !database.ItemExists(dbconfig, TagID) {
                return errors.New("Item Not Yet Registered")
        }

        next_channel := database.NextFinderChannel(dbconfig, FinderOpenID)

        if next_channel < 0 {
                return errors.New("Find Item Limit Reached")
        }

	err := database.FindTag(config.DatabaseConfig, FinderOpenID, next_channel, TagID)
	
	if err == nil{
		err = database.ChangeChannel(dbconfig, FinderOpenID, next_channel)
		if err == nil{
			wechat.SendChannelChangeConfirmation(FinderOpenID, next_channel, config)
		}
	}

	return err
}

func changeChannel(q *structures.Message, config *structures.GlobalConfiguration) error {

	OpenID := q.FromUserName
	Channel , err := strconv.Atoi(q.EventKey)

	err = database.ChangeChannel(config.DatabaseConfig, OpenID, Channel)

        if err == nil{
                err = wechat.SendChannelChangeConfirmation(OpenID, Channel, config)
        }


        strs, err:= queue.Flush(OpenID, Channel, config.RedisAccessInfo)

        for str := len(strs)-1; str >= 0; str-- {
                //log.Println(str)
                time.Sleep(time.Second)
                wechat.Send(strs[str], OpenID, -1, config)
        }

        return err
}



