/* defines basic structures */

package structures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WeChat configuration struct
// Contains:
//	AppId
//	AppSecret
// 	Token - Selfdefined Token
// 	Access - AcessToken used for calling the WeChat API

type Config struct {
	AppId      string
	AppSecret  string
	Token      string
	Access     AccessToken
	Expiration int64

	use chan int
}

// AccessToken struct for receiving JSON response from WeChat API

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// returns a new configuration (should only be called once)

func NewConfig() *Config {

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
	ret.RefreshAccessToken()
	return ret
}

// Refresh the access token in the global configuration struct

func (c *Config) RefreshAccessToken() {

	_, _ = <-c.use

	cur := time.Now().Unix()

	if cur < c.Expiration {
		fmt.Println("AccessTokenStillValid!")
		return
	}

	requrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + c.AppId + "&secret=" + c.AppSecret
	fmt.Println(requrl)
	resp, err := http.Get(requrl)
	if err == nil {
		json.NewDecoder(resp.Body).Decode(&(c.Access))
		c.Expiration = cur + int64(c.Access.ExpiresIn)
		fmt.Println(c.Access)
	}

	c.use <- 1
}

type GlobalConfiguration struct {
        WeChatConfig   *Config
        DatabaseConfig *DatabaseAccessInfo
}


