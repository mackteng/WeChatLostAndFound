package wechat

import(
	"fmt"
	"encoding/json"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)

func PrepareTemplateMessage(OpenID string) string {

	template := structures.TemplateMessage{

		ToUser: OpenID,
		TemplateID : "vC1xeaPgqKfQFSBj6d-8v9YagktAeHE7nDM0IUpaLl8",

	}

	b, err := json.Marshal(template)

	if err != nil {
		
		fmt.Println(err)

	}
	return string(b)

}

func PrepareTextMessage(OpenID string, Message string) string{

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
func PrepareImageMessage(OpenID string, Media string) string{

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

func PrepareVoiceMessage(OpenID string, Media string) string {

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
