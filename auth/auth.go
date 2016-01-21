package auth

import(
	"math/rand"
	"time"
	"crypto/sha1"
	"fmt"
	"net/http"
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
)

var (
	chars              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	defaultRand        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetConfig(r *http.Request, config *structures.GlobalConfiguration) structures.SignPackage {

	timestamp := time.Now().Unix()
	noncestr := createNonceStr(16)
	ticket := config.WeChatInteractor.GetJSApiTicket()	
	
	r.ParseForm()

	url := "http://ec2-52-68-156-216.ap-northeast-1.compute.amazonaws.com/menu?code=" + r.Form["code"][0] + "&state=123" 
		



	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s",
		ticket, noncestr, timestamp, url)
	signature := fmt.Sprintf("%x", sha1.Sum([]byte(str)))

	return structures.SignPackage{
		timestamp, noncestr, signature,
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

