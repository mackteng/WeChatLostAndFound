package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func EntryHandler(r *http.Request, config *structures.GlobalConfiguration) {

	log.Println("Entry Handler Called")
	

	m:= structures.Message{}
	decoder := xml.NewDecoder(r.Body)
        err := decoder.Decode(&m)

	if err != nil {
		log.Fatal("Failed to Parse Message")
	} else {

		switch m.MsgType {
		case "event":
			EventMessageHandler(&m, config)
		default:
			UserMessageHandler(&m, config)
		}
	}
}

func UserMessageHandler(t *structures.Message, config *structures.GlobalConfiguration) {

	log.Println("UserMessageHandlerCalled")
	fmt.Println(t)

}




func EventMessageHandler(q *structures.Message, config *structures.GlobalConfiguration) {

	log.Println("EventMessageHandlerCalled")
	
	if !database.ItemExists(config.DatabaseConfig, q.ScanCodeInfo.ScanResult){
		test := structures.ItemInfo{
				q.ScanCodeInfo.ScanResult,
				"foo",
				"foo",
			}
		
		database.AddItemOwner(config.DatabaseConfig , q.FromUserName, &test)
	}else{

		database.AddItemFinder(config.DatabaseConfig, q.FromUserName, q.ScanCodeInfo.ScanResult)

	}

	fmt.Println(q)

}
