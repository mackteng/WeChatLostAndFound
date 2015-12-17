package controller

import (
	"log"

	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)

func EventHandler(q *structures.Message, config *structures.GlobalConfiguration) error {

	log.Println("EventMessageHandlerCalled", q)

	if !database.ItemExists(config.DatabaseConfig, q.ScanCodeInfo.ScanResult) {
		test := structures.ItemInfo{
			q.ScanCodeInfo.ScanResult,
			"foo",
			"foo",
		}

		database.AddItemOwner(config.DatabaseConfig, q.FromUserName, &test)
	} else {

		database.AddItemFinder(config.DatabaseConfig, q.FromUserName, q.ScanCodeInfo.ScanResult)

	}

	return nil

}
