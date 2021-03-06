package main

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/controller"
	"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bitbucket.org/mack_teng/WeChatLostAndFound/redis"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bitbucket.org/mack_teng/WeChatLostAndFound/wechat"

	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type handle struct {
	config *structures.GlobalConfiguration
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		controller.EntryHandler(r, w, h.config)
		log.Println("DONE")
	} else {
		fmt.Println(r)
		r.ParseForm()
		fmt.Println(r.Body)
		fmt.Fprintf(w, r.Form["echostr"][0])
	}
}

func main() {

	wechat := wechat.NewWeChat()
	database := database.NewDatabase()
	redis := redis.NewRedis()

	w := &structures.GlobalConfiguration{
		WeChatInteractor:   wechat,
		DatabaseInteractor: database,
		RedisInteractor:    redis,
	}

	h := handle{
		config: w,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", &h)
	log.Fatal(http.ListenAndServe(":80", router))
}
