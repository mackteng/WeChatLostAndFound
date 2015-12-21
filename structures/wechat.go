package structures


type Content struct {
	Content string `json:"content"`
}

type MediaID struct {
	MediaID string `json:"media_id"`
}

type WeChatMessage struct {
	ToUser  string  `json:"touser"`
	MsgType string  `json:"msgtype"`
	Text    Content `json:"text"`
	Image   MediaID `json:"image"`
	Voice   MediaID `json:"voice"`
}


