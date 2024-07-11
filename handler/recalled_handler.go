package handler

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/Vingurzhou/wechat-robot/cache"
	"github.com/eatmoreapple/openwechat"
)

func RecalledMatchFunc(m *openwechat.Message) bool {
	cache.Set(m.MsgId, m, 3*time.Minute)
	sender, _ := m.Sender()
	return m.IsRecalled() && sender.NickName != "宝山搭子遍宝山"
}
func RecalledHandler(ctx *openwechat.MessageContext) {
	revokeMsg, _ := ctx.RevokeMsg()
	v := cache.Get(fmt.Sprint(revokeMsg.RevokeMsg.MsgId))
	msg, ok := v.(*openwechat.Message)
	if !ok {
		log.Println(ok)
	}
	senderInGroup, err := msg.SenderInGroup()
	if err != nil {
		log.Println(err)
	}
	if ctx.IsSendBySelf() {
		return
	}
	if err != nil {
		log.Println(err)
	}

	ctx.ReplyText(fmt.Sprintf("%s撤回一条%+v", senderInGroup.NickName, msg.MsgType))

	if msg.IsText() {
		ctx.ReplyText(msg.Content)
	}
	if msg.IsPicture() {
		resp, err := msg.GetPicture()
		if err != nil {
			log.Println(err)
		}
		ctx.ReplyImage(resp.Body)
	}
	if msg.IsVideo() {
		resp, err := msg.GetVideo()
		if err != nil {
			log.Println(err)
		}
		ctx.ReplyVideo(resp.Body)
	}
	if msg.IsEmoticon() {
		var msgContent struct {
			Emoji struct {
				MD5 string `xml:"md5,attr"`
			} `xml:"emoji"`
		}

		err := xml.Unmarshal([]byte(msg.Content), &msgContent)
		if err != nil {
			log.Println(err)
		}
		ctx.ReplyEmoticon(msgContent.Emoji.MD5, nil)
	}
}
