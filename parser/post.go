package parser

import(
	"fmt"
	"net/http"
	"encoding/xml"
	
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)


func Parsepost(r *http.Request) (t structures.Message, err error) {


	decoder := xml.NewDecoder(r.Body)
	err = decoder.Decode(&t)	
	if err== nil{
		fmt.Println(t)
	}
	return
}
