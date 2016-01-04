package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bitbucket.org/mack_teng/WeChatLostAndFound/wechat"
	"log"
)

func TextMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {
	
	log.Println("TextMessageHandler")	
	OpenID := m.FromUserName
	SendToID, Channel, err := database.FindCorrespondingUser(config.DatabaseConfig, OpenID)
	
	if err == nil {
		wechat.SendMessage(m.Content, SendToID, Channel, config)
	}
	
	return err
}

func ImageMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {
	log.Println("ImageMessageHandler", m)
	return nil

}

func VoiceMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {
	log.Println("VoiceMessageHandler", m)
	return nil

}

func VideoMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {
	log.Println("VideoMessageHandler", m)
	return nil

}

func LocationMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {
	log.Println("LocationMessageHandler", m)
	return nil

}

