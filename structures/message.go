package structures

import "time"

type Message struct {
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

	Event      string `xml:"Event"`
	EventKey   string `xml:"EventKey"`

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


func (c *Message) GetMsgType() string {
	if c.MsgType == "event"{
		return c.Event	
	} else{
		return c.MsgType
	}
}

