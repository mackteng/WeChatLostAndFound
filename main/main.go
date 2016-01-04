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
	request_table *structures.Set
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {		
		controller.EntryHandler(r, w, h.request_table,  h.config)
		log.Println("DONE")
	} else {
		fmt.Println(r)
		r.ParseForm()
		fmt.Println(r.Body)
		fmt.Fprintf(w, r.Form["echostr"][0])
	}
}

func main() {

	w := structures.InitGlobalConfig()

	h := handle{
		config: w,
		request_table : &structures.Set{
					make(map[string] bool),
				},
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", &h)
	log.Fatal(http.ListenAndServe(":80", router))
}
