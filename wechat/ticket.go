package wechat

import (
	"encoding/json"
	"net/http"
	"time"
)

const accesstokenurl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + "wx97b3ede422c4956e" + "&secret=" + "d4624c36b6795d1d99dcf0547af5443d"

func (c *AccessTokenServer) refreshToken() {
	c.lock.RLock()
	defer c.lock.RUnlock()
	cur := time.Now().Unix()

	if cur < c.Expiration {
		return
	}

	resp, err := http.Get(accesstokenurl)
	if err == nil {
		json.NewDecoder(resp.Body).Decode(&(c.CachedAccessToken))
		c.Expiration = cur + int64(c.CachedAccessToken.ExpiresIn)
	}

}

func (c *AccessTokenServer) getAccessToken() string {
	c.refreshToken()
	return c.CachedAccessToken.AccessToken
}

func (c *JSTicketServer) refreshTicket(AccessToken string) {
	c.lock.RLock()
        defer c.lock.RUnlock()
	cur := time.Now().Unix()

	if cur < c.Expiration {
		return
	}

	requrl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + AccessToken + "&type=jsapi"
	resp, err := http.Get(requrl)

	if err != nil {
		return
	}

	json.NewDecoder(resp.Body).Decode(&c.CachedJSTicket)
	c.Expiration = cur + int64(c.CachedJSTicket.ExpiresIn)
}

func (c *JSTicketServer) getJSApiTicket(AccessToken string) string {
	c.refreshTicket(AccessToken)
	return c.CachedJSTicket.Ticket
}
