package wechat

import(
	"fmt"
	"encoding/json"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)


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
