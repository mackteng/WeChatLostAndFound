package wechat

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	AppId      string
	AppSecret  string
	Token      string
	Access     AccessToken
	Expiration int64

	use chan int
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WeChat struct {
	AccessConfig *Config
}

const (
	USER     = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
	TEMPLATE = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="

	ALERT_MESSAGE_TEMPLATE_ID = "vC1xeaPgqKfQFSBj6d-8v9YagktAeHE7nDM0IUpaLl8" 
)

func (c *Config) refreshAccessToken() {

	log.Println("Refreshring Access Token")
	cur := time.Now().Unix()

	if cur < c.Expiration {
		fmt.Println("AccessTokenStillValid!")
		return
	}

	_, _ = <-c.use

	requrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + c.AppId + "&secret=" + c.AppSecret
	resp, err := http.Get(requrl)
	if err == nil {
		json.NewDecoder(resp.Body).Decode(&(c.Access))
		c.Expiration = cur + int64(c.Access.ExpiresIn)
		fmt.Println(c.Access)
	}

	c.use <- 1
}

func (c *Config) GetAccessToken() string {
	log.Println("Getting Access Token")
	c.refreshAccessToken()
	return c.Access.AccessToken

}

func (w *WeChat) GetAccessToken() string {

	return w.AccessConfig.GetAccessToken()

}

func (w *WeChat) GetJSApiTicket() string {

	response := struct {
		Code      int    `json:"errcode"`
		Msg       string `json:"errmsg"`
		Ticket    string `json:"ticket"`
		ExpiresIn int    `json:"expires_in"`
	}{}

	requrl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + w.GetAccessToken() + "&type=jsapi"
	resp, err := http.Get(requrl)
	
	if err != nil {
		return ""
	}
	
	json.NewDecoder(resp.Body).Decode(&response)

	return response.Ticket

}

func NewWeChat() *WeChat {

	ret := &Config{
		AppId:     "wx97b3ede422c4956e",
		AppSecret: "d4624c36b6795d1d99dcf0547af5443d",
		Token:     "macktengmackteng",
		Access: AccessToken{
			AccessToken: "",
			ExpiresIn:   0,
		},
	}
	ret.use = make(chan int, 1)
	ret.use <- 1
	_ = ret.GetAccessToken()
	return &WeChat{
		AccessConfig: ret,
	}
}

func (WeChat *WeChat) SendSystemMessage(OpenID string, SystemID string, Config *structures.GlobalConfiguration) error {

	Payload := prepareTextMessage(OpenID, SystemID)
	return send(Payload, USER, Config)
}

func (WeChat *WeChat) SendTemplateMessage(OpenID, TemplateID string, Config *structures.GlobalConfiguration) error {

	Payload := prepareTemplateMessage(OpenID, TemplateID)
	return send(Payload, TEMPLATE, Config)
}

func (WeChat *WeChat) SendForwardMessage(Msg string, OpenID string, Channel int, Config *structures.GlobalConfiguration) error {

	cur_channel, _ := Config.DatabaseInteractor.CurrentChannel(OpenID)
	Payload := prepareTextMessage(OpenID, Msg)
        Config.RedisInteractor.AddMessageToQueue(OpenID, Channel, Payload)
	if cur_channel == Channel {
                        return send(Payload, USER, Config)
        } else {
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
