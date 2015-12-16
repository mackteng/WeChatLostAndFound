package parser

import (
	"bitbucket.org/mack_teng/WeChatLostAndFound/structures"
	"encoding/xml"
	"net/http"
	"time"
)

func ParsePost(r *http.Request) (t structures.Message, err error) {

	decoder := xml.NewDecoder(r.Body)

	// universal properties
	var ToUserName string
	var FromUserName string
	var CreateTime time.Duration
	var MsgType string

	// Text Message
	var Content string
	var MsgId int


	// All Types of Events
	var Event string
	var EventKey string
	var ScanResult string

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			name := se.Name.Local
			if name == "ToUserName" {
				decoder.DecodeElement(&ToUserName, &se)
			} else if name == "FromUserName" {
				decoder.DecodeElement(&FromUserName, &se)
			} else if name == "CreateTime" {
				decoder.DecodeElement(&CreateTime, &se)
			} else if name == "MsgType" {
				decoder.DecodeElement(&MsgType, &se)
			} else if name == "Content" {
				decoder.DecodeElement(&Content, &se)
			} else if name == "MsgId" {
				decoder.DecodeElement(&MsgId, &se)
			} else if name == "Event" {
				decoder.DecodeElement(&Event, &se)
			} else if name == "EventKey" {
				decoder.DecodeElement(&EventKey, &se)
			} else if name == "ScanResult" {
				decoder.DecodeElement(&ScanResult, &se)
			}
		}

	}

	if MsgType != "event" {
		tmp := &structures.UserMessage{}
		tmp.ToUserName = ToUserName
		tmp.FromUserName = FromUserName
		tmp.CreateTime = CreateTime
		tmp.MsgType = MsgType
		tmp.Content = Content
		tmp.MsgId = MsgId
		return tmp, nil
	} else {
		tmp := &structures.EventMessage{}
		tmp.ToUserName = ToUserName
		tmp.FromUserName = FromUserName
		tmp.CreateTime = CreateTime
		tmp.MsgType = MsgType
		tmp.Event = Event
		tmp.EventKey = EventKey
		tmp.ScanResult = ScanResult
		return tmp, nil
	}
	return
}
