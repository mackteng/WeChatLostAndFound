package main

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/controller"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"

	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type handle struct {
	config *structures.GlobalConfiguration
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {		
		controller.EntryHandler(r, w, h.config)
		log.Println("DONE")
	} else {
		r.ParseForm()
		fmt.Println(r.Form["echoStr"][0])
		fmt.Fprintf(w, r.Form["echoStr"][0])
	}
}

func main() {

	w := structures.InitGlobalConfig()

	h := handle{
		config: w,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", &h)
	log.Fatal(http.ListenAndServe(":80", router))
}
