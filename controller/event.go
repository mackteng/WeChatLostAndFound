package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures/sysmsg"
	"log"
)

var entryhandlers = map[string]handler{
	"RegisterTag": RegisterTag,
	"FindTag":     FindTag,
}

func EventHandler(q *structures.Message, config *structures.GlobalConfiguration) error {

	log.Println("EventMessageHandlerCalled")

	if q.Event == "subscribe" {
		return Subscribe(q, config)
	} else {
		if val, ok := entryhandlers[q.EventKey]; ok{
			return val(q, config)
		} else {
			return nil
		}
	}

}


func Subscribe(q *structures.Message, config *structures.GlobalConfiguration) error {

	log.Println("Subscribe", q.FromUserName)
	err := config.DatabaseInteractor.AddUser(q.FromUserName)
	
	if err!=nil {
		return err
	}	
	return config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.WELCOME_MESSAGE, config)
}


func RegisterTag(q *structures.Message, config *structures.GlobalConfiguration) error {

	err := config.DatabaseInteractor.RegisterTag(q.FromUserName, q.ScanCodeInfo.ScanResult, q.ItemInfo)
	if err != nil {
		config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.REGISTER_FAIL, config)
		return err
	}

	err = config.DatabaseInteractor.ChangeActiveTag(q.FromUserName, q.ScanCodeInfo.ScanResult);
	if err != nil {
		config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.REGISTER_FAIL, config)
		return err
	}

	config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.REGISTER_SUCCESS, config)
	return nil
}

func FindTag(q *structures.Message, config *structures.GlobalConfiguration) error {

	err := config.DatabaseInteractor.FindTag(q.FromUserName, q.ScanCodeInfo.ScanResult)
	if err != nil {
		config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.FIND_FAIL, config)
		return err
	}

	err = config.DatabaseInteractor.ChangeActiveTag(q.FromUserName, q.ScanCodeInfo.ScanResult);
        if err != nil {
                config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.FIND_FAIL, config)
                return err
        }


	config.WeChatInteractor.SendSystemMessage(q.FromUserName, sysmsg.FIND_SUCCESS, config)
	return nil
}


















/*
func changeChannel(q *structures.Message, config *structures.GlobalConfiguration) error {

	OpenID := q.FromUserName
	Channel, err := strconv.Atoi(q.EventKey)

	err = config.DatabaseInteractor.ChangeChannel(OpenID, Channel)

	if err == nil {
		err = config.WeChatInteractor.SendSystemMessage(OpenID, sysmsg.CHANNEL_CHANGE + q.EventKey, config)
	} else {
		config.WeChatInteractor.SendSystemMessage(OpenID, sysmsg.CHANNEL_CHANGE_FAIL, config)
		return err
	}

	strs, err := config.RedisInteractor.GetMessagesFromQueue(OpenID, Channel)
	config.WeChatInteractor.SendBulkForwardMessages(strs, OpenID, config)
	return err
}*/
