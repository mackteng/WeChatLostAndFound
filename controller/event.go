package controller

import (
	"log"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)

var entryhandlers = map[string]handler{
	"RegisterTag":RegisterTag,
	"FindTag" : FindTag,
	
	"1": OwnerChannel1,	
	"2": OwnerChannel2,	
	"3": OwnerChannel3,	
	"4": OwnerChannel4,	
	"5": OwnerChannel5,

	"6": FinderChannel1,
	"7": FinderChannel2,
	"8": FinderChannel3,
	"9": FinderChannel4,
	"10": FinderChannel5,
	
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

	err := database.RegisterTag(config.DatabaseConfig, q.FromUserName, &test)
	return err
}


func FindTag(q *structures.Message, config *structures.GlobalConfiguration) error{

	log.Println("FindTag Called", q)
	err := database.FindTag(config.DatabaseConfig, q.FromUserName, q.ScanCodeInfo.ScanResult)
	return err
}


func OwnerChannel1(q *structures.Message, config *structures.GlobalConfiguration) error{

	err:= database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 1)
	return err
}

func OwnerChannel2(q *structures.Message, config *structures.GlobalConfiguration) error{

	err:= database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 2)
	return err
}

func OwnerChannel3(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 3)
	return err
}

func OwnerChannel4(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 4)
	return err
}

func OwnerChannel5(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 5)
	return err
}


func FinderChannel1(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 6)
	return err
}


func FinderChannel2(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 7)
	return err
}


func FinderChannel3(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 8)
	return err
}


func FinderChannel4(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 9)
	return err
}


func FinderChannel5(q *structures.Message, config *structures.GlobalConfiguration) error{

	err := database.ChangeChannel(config.DatabaseConfig, q.FromUserName, 10)
	return err
}
















