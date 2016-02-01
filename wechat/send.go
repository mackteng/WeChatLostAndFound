package wechat

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"bytes"
	"log"
	"net/http"
)

func post(Payload string, AccessUrl string, config *structures.GlobalConfiguration) error {

	log.Println("send", Payload)

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
