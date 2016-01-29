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
	return nil
}

func send(Payload string, AccessUrl string, config *structures.GlobalConfiguration) error {
	return post(Payload, AccessUrl, config)
}
