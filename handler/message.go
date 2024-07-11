package handler

import (
	"strings"

	"github.com/eatmoreapple/openwechat"
)

type WrapperMessage struct {
	*openwechat.Message
}

func NewWrapperMessage(m *openwechat.Message) *WrapperMessage {
	return &WrapperMessage{m}
}
func (m *WrapperMessage) WrapperIsJoinGroup() bool {
	return strings.Contains(m.Content, "invited") || strings.Contains(m.Content, "joined") || m.IsJoinGroup()
}
