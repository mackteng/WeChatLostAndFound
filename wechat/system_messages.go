package wechat

import(
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"strconv"
)


func SendChannelChangeConfirmation(OpenID string, channel int, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "----       頻道" + strconv.Itoa(channel) + "       ----")
	Send(Payload, OpenID, -1, config)
	return nil
}

func SendItemRegisteredConfirmation(OpenID string, config *structures.GlobalConfiguration) error {
	
	Payload := PrepareTextMessage(OpenID, "---- 	  物品已成功註冊  ---- ")
	return Send(Payload, OpenID, -1, config)
}


func SendItemAlreadyRegistered(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 該物品已經被註冊! ----")
	Send(Payload, OpenID, -1, config)
	return nil
}

func SendItemLimit(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 註冊物品已達上限! ----")
	Send(Payload, OpenID, -1, config)
	return nil
}

func SendItemNotYetRegistered(OpenID string, config *structures.GlobalConfiguration) error {

	Payload := PrepareTextMessage(OpenID, "---- 該物品還未被註冊! ----")
	Send(Payload, OpenID, -1, config)
	return nil
}


