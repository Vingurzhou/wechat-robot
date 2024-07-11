package handler

import (
	"github.com/eatmoreapple/openwechat"
)

func JoinMatchFunc(m *openwechat.Message) bool {
	return NewWrapperMessage(m).WrapperIsJoinGroup()
}
func JoinHandler(ctx *openwechat.MessageContext) {
	ctx.ReplyEmoticon("4d10f4c2103deb03f7761ae10e22d09b", nil)
}
