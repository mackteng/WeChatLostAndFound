package wechat

import(
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"net/http"
	//"encoding/xml"	
	"fmt"
	"log"
)

type handler func(*structures.Message, http.ResponseWriter) error

var handlers = map[string]handler{

        "text":     EchoTextMessage,
        //"image":    EchoImageMessage,
        //"voice":    EchoVoiceMessage,
        //"video":    EchoVideoMessage,
        //"location": EchoLocationMessage,
        //"event": EchoEvent,
}

func Echo(m *structures.Message, w http.ResponseWriter) {

	log.Println("Echo Entry Handler")	
	handlers[m.GetMsgType()](m, w);
}


const(

	replyHeadera = "<ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime>"
	replyText   = "<xml>%s<MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>"

)


func replyHeader(w *structures.Message) string {
	return fmt.Sprintf(replyHeadera, w.FromUserName, w.ToUserName, w.CreateTime)
}

func ReplyText(m *structures.Message) string {
	msg := fmt.Sprintf(replyText, replyHeader(m), m.Content)
	return msg
}


func EchoTextMessage(m *structures.Message, w http.ResponseWriter) error{

	msg:= ReplyText(m)
	log.Println(msg)
	w.Write([]byte(msg))	

	log.Println("Echo Done")

	return nil

}
