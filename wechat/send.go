package wechat

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	//"bitbucket.org/mack_teng/WeChatLostAndFound/redis"
	//"bitbucket.org/mack_teng/WeChatLostAndFound/database"
	"bytes"
	"fmt"
	"net/http"
	//"io/ioutil"
)

func post(Payload string, AccessUrl string, config *structures.GlobalConfiguration) error {

	fmt.Println("Sending Message: ", Payload)

	url := AccessUrl + config.WeChatInteractor.GetAccessToken()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(Payload)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	return nil

}

func send(Payload string, AccessUrl string, config *structures.GlobalConfiguration) error {
/*
	var err error
	if Channel != -1 {
		config.RedisInteractor.AddMessageToQueue(OpenID, Channel, Payload)
		cur_channel, _ := config.DatabaseInteractor.CurrentChannel(OpenID)

		if cur_channel == Channel {
			post(Payload, AccessUrl, config)
		}

	} else {
		err = post(Payload, AccessUrl, config)
	}
*/

	
	return post(Payload, AccessUrl, config)
}
