package structures

import "encoding/xml"

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

type EchoTextMessage struct{
	XMLName      xml.Name `xml:"xml"`
	ToUserName  string  `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:Content"`
}
