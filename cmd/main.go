package main

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/parser"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GlobalConfiguration struct {
	WeChatConfig   *structures.Config
	DatabaseConfig *structures.DatabaseAccessInfo
}

func (h *GlobalConfiguration) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(parser.ParsePost(r))
	fmt.Fprintf(w, "hello")
}

func main() {

	configuration := structures.NewConfig()
	configuration.RefreshAccessToken()
	database := structures.NewDatabase()

	w := GlobalConfiguration{
		WeChatConfig:   configuration,
		DatabaseConfig: database,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", &w)

	log.Fatal(http.ListenAndServe(":80", router))
}
