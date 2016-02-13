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
		items,  _ := config.DatabaseInteractor.GetAllOwnedItems(signPackage.OpenID)	
		ActiveTag,  _ := config.DatabaseInteractor.GetActiveTag(signPackage.OpenID)	
		t, _ := template.ParseFiles("/home/ubuntu/work/src/bitbucket.org/mack_teng/WeChatLostAndFound/menu/template.html")
		
		err := t.Execute(w, struct{
					Auth interface{}
					Items interface{}
					Active string
				}{
					signPackage,
					items,
					ActiveTag,
				})

		if err!=nil {
			log.Println(err)
		}

	}
}
