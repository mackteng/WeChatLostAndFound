package parser

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"encoding/xml"
	"net/http"
)

func ParsePost(r *http.Request) (t structures.Message, err error) {

	decoder := xml.NewDecoder(r.Body)
	err = decoder.Decode(&t)
	return 
}
