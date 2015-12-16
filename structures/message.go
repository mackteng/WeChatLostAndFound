package structures

import "time"

type Message interface {
	GetMsgType() string
}

type UserMessage struct {
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	MsgId   int    `xml:"MsgId"`

	
	Content string `xml:"Content"` // text message
	
	MediaId string // voice and video message id
	Format string  // voice message format
	
	ThumbMediaId string // video message thumbnail
	
	Location_X string // location x
	Location_Y string // location y
	Scale int // Scale
	Label string // Map
}

type EventMessage struct {
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`

	Event      string `xml:"Event"`
	EventKey   string `xml:"EventKey"`

	ScanResult string
	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
}

type ScanCodeInfo struct{
        ScanType string
        ScanResult string
}

type ItemInfo struct{
	TagID string
	Name string
	Description string
}


func (c *EventMessage) GetMsgType() string {
	return c.MsgType
}

func (c *UserMessage) GetMsgType() string {
	return c.MsgType
}
