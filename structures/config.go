/* defines basic structures */

package structures

import(
	"fmt"
	"net/http"
	"encoding/json"
)

// global configuration struct 
// Contains:
//	AppId
//	AppSecret
// 	Token - Selfdefined Token
// 	Access - AcessToken used for calling the WeChat API

type config struct{
	AppId string
	AppSecret string
	Token string
	Access AccessToken 
}

// AccessToken struct for receiving JSON response from WeChat API

type AccessToken struct {
        AccessToken string `json:"access_token"`
        ExpiresIn int `json:"expires_in"`
}


// returns a new configuration (should only be called once)

func NewConfig() *config{

	ret := &config{
		AppId: "wx97b3ede422c4956e",
		AppSecret: "d4624c36b6795d1d99dcf0547af5443d",
		Token: "macktengmackteng",
		Access: AccessToken{
				AccessToken: "",
				ExpiresIn: 0,
			     },
	}
	
	ret.RefreshAccessToken()
	return ret
}

// Refresh the access token in the global configuration struct

func (c *config) RefreshAccessToken() {
	requrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + c.AppId + "&secret=" + c.AppSecret
	fmt.Println(requrl)
	resp, err := http.Get(requrl)
	if err == nil{
		json.NewDecoder(resp.Body).Decode(&(c.Access))
		fmt.Println(c.Access)
	}
}


