package menu

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/auth"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"html/template"
	"net/http"
	"log"
)

func MenuHandler(config *structures.GlobalConfiguration) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Menu")
		signPackage := auth.GetConfig(r, config)
		items, _ := config.DatabaseInteractor.GetAllOwnedItems(signPackage.OpenID)		
		t, _ := template.ParseFiles("/home/ubuntu/work/src/bitbucket.org/mack_teng/WeChatLostAndFound/menu/template.html")
		log.Println(items)
		err := t.Execute(w, struct{
					Auth interface{}
					Items interface{}
				}{
					signPackage,
					items,
				})


		log.Println(err)

	}
}
