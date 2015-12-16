package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/parser"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"fmt"
	"log"
	"net/http"
)

func EntryHandler(r *http.Request, config *structures.GlobalConfiguration) {

	log.Println("Entry Handler Called")
	m, err := parser.ParsePost(r)

	if err != nil {
		log.Fatal("Failed to Parse Message")
	} else {

		switch m.(type) {
		case *structures.UserMessage:
			textmessage := m.(*structures.UserMessage)
			UserMessageHandler(textmessage, config)
		case *structures.EventMessage:
			qrcodemessage := m.(*structures.EventMessage)
			EventHandler(qrcodemessage, config)
		}
	}
}

func UserMessageHandler(t *structures.UserMessage, config *structures.GlobalConfiguration) {

	log.Println("UserMessageHandlerCalled")
	fmt.Println(t)

}




func EventHandler(q *structures.EventMessage, config *structures.GlobalConfiguration) {

	log.Println("EventMessageHandlerCalled")
	
	if !database.ItemExists(config.DatabaseConfig, q.ScanResult){
		test := structures.ItemInfo{
				q.ScanResult,
				"foo",
				"foo",
			}
		
		database.AddItemOwner(config.DatabaseConfig , q.FromUserName, &test)
	}else{

		database.AddItemFinder(config.DatabaseConfig, q.FromUserName, q.ScanResult)

	}

	fmt.Println(q)

}
