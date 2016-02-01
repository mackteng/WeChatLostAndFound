package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"log"
)

func TextMessageHandler(m *structures.Message, config *structures.GlobalConfiguration) error {

	log.Println("TextMessageHandler")
	OpenID := m.FromUserName
	ActiveTag, SendToID, err := config.DatabaseInteractor.FindCorrespondingUser(OpenID)
	log.Println(SendToID, ActiveTag)
	if err == nil {
		return config.WeChatInteractor.SendForwardMessage(m.Content, SendToID, ActiveTag, config)
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
