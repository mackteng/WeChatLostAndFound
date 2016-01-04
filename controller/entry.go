package controller

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	//"bitbucket.org/mack_teng/WeChatLostAndFound/wechat"
	//"time"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"
)

type handler func(*structures.Message, *structures.GlobalConfiguration) error

var handlers = map[string]handler{

	"text":     TextMessageHandler,
	"image":    ImageMessageHandler,
	"voice":    VoiceMessageHandler,
	"video":    VideoMessageHandler,
	"location": LocationMessageHandler,

	"event": EventHandler,
}

func EntryHandler(r *http.Request, w http.ResponseWriter, t *structures.Set, config *structures.GlobalConfiguration) {

	log.Println("Entry Handler Called")

	m := structures.Message{}
	decoder := xml.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	

	if err != nil {
		log.Fatal("Failed to Parse Message")
	} else {
		w.Write([]byte("success"))
		if f,ok := w.(http.Flusher); ok {
			f.Flush()
		}

		if found := t.Add(m.FromUserName + strconv.FormatInt(m.CreateTime,10) + strconv.Itoa(m.MsgId)); found {

			handlers[m.GetMsgType()](&m, config)

		}
	}
}
