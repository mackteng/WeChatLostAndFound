package structures

import "time"

type Message interface {
	GetMsgType() string
}

type TextMessage struct {
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`

	Content string `xml:"Content"`
	MsgId   int    `xml:"MsgId"`
}

type QRCodeMessage struct {
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`

	Event      string `xml:"Event"`
	EventKey   string `xml:"EventKey"`
	ScanResult string
}

func (c *QRCodeMessage) GetMsgType() string {
	return c.MsgType
}

func (c *TextMessage) GetMsgType() string {
	return c.MsgType
}
