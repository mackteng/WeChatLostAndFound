package menu

import(
	"net/http"
	"log"
	"html/template"
	"bitbucket.org/mack_teng/WeChatLostAndFound/auth"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)



func MenuHandler(config *structures.GlobalConfiguration) func(w http.ResponseWriter, r *http.Request)  {

	return func(w http.ResponseWriter, r * http.Request) {
		signPackage := auth.GetConfig(r, config)
		log.Println(signPackage) 
		t, _ := template.ParseFiles("/home/ubuntu/work/src/bitbucket.org/mack_teng/WeChatLostAndFound/menu/menu.html")		
		t.Execute(w, signPackage)
	}
}


