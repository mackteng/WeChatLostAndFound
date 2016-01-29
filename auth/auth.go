package auth

import(
	"math/rand"
	"time"
	"crypto/sha1"
	"fmt"
	"net/http"
	"encoding/json"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)

var (
	chars              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	defaultRand        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetConfig(r *http.Request, config *structures.GlobalConfiguration) structures.SignPackage {
	
	r.ParseForm()
	


	timestamp := time.Now().Unix()
	noncestr := createNonceStr(16)
	ticket := config.WeChatInteractor.GetJSApiTicket()	
	openid := getOpenIDFromCode(r.Form["code"][0])	


	url := "http://" + r.Host + r.URL.Path + "?code=" + r.Form["code"][0] + "&state=123" 
	fmt.Println(url)



	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s",
		ticket, noncestr, timestamp, url)
	signature := fmt.Sprintf("%x", sha1.Sum([]byte(str)))

	return structures.SignPackage{
		openid, timestamp, noncestr, signature, 
	}

}


func createNonceStr(length int) string {
	var str string
	for i := 0; i < length; i++ {
		tmpI := defaultRand.Intn(len(chars) - 1)
		str += chars[tmpI : tmpI+1]
	}
	return str
}

func getOpenIDFromCode(code string) string {


        response := struct {
                OpenID string `json:"openid"`
        }{}

        requrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx97b3ede422c4956e&secret=d4624c36b6795d1d99dcf0547af5443d&code=" + code + "&grant_type=authorization_code";

        resp, err := http.Get(requrl)

        if err != nil{
                return ""
        }

        json.NewDecoder(resp.Body).Decode(&response)

        return response.OpenID
}

