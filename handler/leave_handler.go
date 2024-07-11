package handler

import (
	"fmt"
	"log"

	"github.com/Vingurzhou/wechat-robot/cache"
	"github.com/eatmoreapple/openwechat"
)

func LeaveMatchFunc(m *openwechat.Message) bool {
	return true
}
func LeaveHandler(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	group, ok := sender.AsGroup()
	if !ok {
		log.Println(ok)
		return
	}
	v1, ok := cache.Get(group.UserName).(*openwechat.Group)
	if !ok {
		cache.Set(group.UserName, group, 0)
		log.Println("set success")
		return
	}
	for _, ele := range v1.MemberList {
		_, err := group.SearchMemberByUsername(ele.UserName)
		if err != nil {
			ctx.ReplyText(fmt.Sprintf("%s退群", ele.NickName))
		}
	}
}
