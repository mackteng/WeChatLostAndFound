package wechat

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"time"
	"log"
	"sync"
)

const (
	USER                      = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
	TEMPLATE                  = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
	ALERT_MESSAGE_TEMPLATE_ID = "vC1xeaPgqKfQFSBj6d-8v9YagktAeHE7nDM0IUpaLl8"
)

type WeChat struct {
	AppId     string
	AppSecret string
	Token     string
	access    AccessTokenServer
	ticket    JSTicketServer
}
type AccessTokenServer struct {
	CachedAccessToken AccessToken
	lock		  sync.RWMutex
	Expiration        int64
}

type JSTicketServer struct {
	CachedJSTicket JSTicket
	lock           sync.RWMutex
	Expiration     int64
}

type JSTicket struct {
	Code      int    `json:"errcode"`
	Msg       string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewWeChat() *WeChat {

	ret := &WeChat{
		AppId:     "wx97b3ede422c4956e",
		AppSecret: "d4624c36b6795d1d99dcf0547af5443d",
		Token:     "macktengmackteng",
	}
	log.Println("AccessToken", ret.GetAccessToken())
	log.Println("JSTicket", ret.GetJSApiTicket())
	return ret
}

func (WeChat *WeChat) GetJSApiTicket() string {
	return WeChat.ticket.getJSApiTicket(WeChat.GetAccessToken())
}

func (WeChat *WeChat) GetAccessToken() string {
	return WeChat.access.getAccessToken()
}
func (WeChat *WeChat) SendSystemMessage(OpenID string, SystemID string, Config *structures.GlobalConfiguration) error {

	Payload := prepareTextMessage(OpenID, SystemID)
	return send(Payload, USER, Config)
}

func (WeChat *WeChat) SendTemplateMessage(OpenID, TemplateID string, Config *structures.GlobalConfiguration) error {

	Payload := prepareTemplateMessage(OpenID, TemplateID)
	return send(Payload, TEMPLATE, Config)
}

func (WeChat *WeChat) SendForwardMessage(Msg string, OpenID string, TagID string, Config *structures.GlobalConfiguration) error {

	Payload := prepareTextMessage(OpenID, Msg)
	ActiveTag, err := Config.DatabaseInteractor.GetActiveTag(OpenID)

	if err != nil {
		return err
	}

	if TagID == ActiveTag {
		return send(Payload, USER, Config)
	} else {
		Config.RedisInteractor.AddMessageToQueue(OpenID, TagID, Payload)
		return WeChat.SendTemplateMessage(OpenID, ALERT_MESSAGE_TEMPLATE_ID, Config)
	}
}

func (WeChat *WeChat) SendBulkForwardMessages(strs []string, OpenID string, Config *structures.GlobalConfiguration) error {

	for str := len(strs) - 1; str >= 0; str-- {
		time.Sleep(time.Second)
		send(strs[str], USER, Config)
	}
	return nil
}
