package main

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/parser"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	configuration := structures.NewConfig()
	configuration.RefreshAccessToken()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(":80", router))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(parser.ParsePost(r))
	fmt.Fprintf(w, "hello")
}
