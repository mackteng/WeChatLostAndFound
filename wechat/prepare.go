package wechat

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"encoding/json"
	"fmt"
)

func prepareTemplateMessage(OpenID string, TemplateID string) string {

	template := structures.TemplateMessage{

		ToUser:     OpenID,
		TemplateID: TemplateID,
	}

	b, err := json.Marshal(template)

	if err != nil {

		fmt.Println(err)

	}
	return string(b)

}

func prepareTextMessage(OpenID string, Message string) string {

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

/*
func prepareImageMessage(OpenID string, Media string) string{

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

func prepareVoiceMessage(OpenID string, Media string) string {

        config := structures.WeChatMessage{
                ToUser:  OpenID,
                MsgType: "voice",
                Text: structrues.MediaID{
                        MediaID: Media,
                },
        }

        b, err := json.Marshal(config)

        if err != nil {
                fmt.Println(err)
        }

        return string(b)
}*/
