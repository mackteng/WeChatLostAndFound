package controller

import (
        "bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"log"
)


func TextMessageHandler (m *structures.Message, config *structures.GlobalConfiguration) error{
	log.Println("TextMessageHandler", m)
	return nil

}


func ImageMessageHandler (m *structures.Message, config *structures.GlobalConfiguration) error{
	log.Println("ImageMessageHandler", m)
	return nil

}


func VoiceMessageHandler (m *structures.Message, config *structures.GlobalConfiguration) error{
	log.Println("VoiceMessageHandler", m)
	return nil

}


func VideoMessageHandler (m *structures.Message, config *structures.GlobalConfiguration) error{
	log.Println("VideoMessageHandler", m)
	return nil

}


func LocationMessageHandler (m *structures.Message, config *structures.GlobalConfiguration) error{
	log.Println("LocationMessageHandler", m)
	return nil

}


