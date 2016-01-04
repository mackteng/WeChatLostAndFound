package wechat

import(
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"strconv"
	"time"
)


const (

	USER = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
	TEMPLATE = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="

)


func SendMessage(Msg string, OpenID string, Channel int, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, Msg) 
	return Send(Payload, USER, OpenID, Channel, config)

}



func SendBulk( strs []string, OpenID string, config *structures.GlobalConfiguration) {

	for str := len(strs)-1; str >= 0; str-- {
                time.Sleep(time.Second)
		Send(strs[str], USER, OpenID, -1, config)
        }
} 



func Alert(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTemplateMessage(OpenID)
	return Send(Payload, TEMPLATE, "", -1, config)

}


func SendChannelChangeConfirmation(OpenID string, channel int, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "----       頻道" + strconv.Itoa(channel) + "       ----")
	return Send(Payload, USER, OpenID, -1, config)
}

func SendItemRegisteredConfirmation(OpenID string, config *structures.GlobalConfiguration) error {
	
	Payload := PrepareTextMessage(OpenID, "---- 	  物品已成功註冊  ---- ")
	return Send(Payload, USER, OpenID, -1, config)
}


func SendItemAlreadyRegistered(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 該物品已經被註冊! ----")
	return Send(Payload, USER, OpenID, -1, config)
}

func SendItemLimit(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 註冊物品已達上限! ----")
	return Send(Payload, USER, OpenID, -1, config)
}

func SendItemNotYetRegistered(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 該物品還未被註冊! ----")
	return Send(Payload, USER,OpenID, -1, config)
}


