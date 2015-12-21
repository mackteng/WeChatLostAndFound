package wechat

import(
	"fmt"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures
)

func SendTextMessage(OpenID string, Message string) string {

	config := structures.WeChatMessage{
		ToUser:  OpenID,
		MsgType: "text",
		Text: structures.Content{
			Content: Message,
		},
	}

	b, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func SendImageMessage(OpenID string, Media string) string{
	
	config := structures.WeChatMessage{
			ToUser: OpenID,
			MsgType: "image",
			Image: structures.MediaID{
				MediaID: Media,
			},
	}
	

	b, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func SendVoiceMessage(OpenID string, Message string) string {

	config := structures.WeChatMessage{
		ToUser:  OpenID,
		MsgType: "voice",
		Text: structrues.Content{
			Content: Message,
		},
	}

	b, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

