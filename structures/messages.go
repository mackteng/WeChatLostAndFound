package structures

type Message struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MsgId        int    `xml:"MsgId"`

	Content string `xml:"Content"` // text message

	MediaId string // voice and video message id
	Format  string // voice message format

	ThumbMediaId string // video message thumbnail

	Location_X string // location x
	Location_Y string // location y
	Scale      int    // Scale
	Label      string // Map

	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`

	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
	ItemInfo     ItemInfo     `xml:"ItemInfo"`
}

type ScanCodeInfo struct {
	ScanType   string
	ScanResult string
}

type SignPackage struct {
	OpenID    string `json:"openid"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

type ItemInfo struct {
	Name        string
	Description string
	TagID string
	FinderID string
}

func (c *Message) GetMsgType() string {
	return c.MsgType
}

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

type TemplateMessage struct {
	ToUser     string `json:"touser"`
	TemplateID string `json:"template_id"`
}
