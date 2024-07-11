package handler

import (
	"github.com/eatmoreapple/openwechat"
)

func NewMessageHandler() openwechat.MessageHandler {

	dispatcher := openwechat.NewMessageMatchDispatcher()

	dispatcher.RegisterHandler(TestMatchFunc, TestHandler)

	dispatcher.RegisterHandler(JoinMatchFunc, JoinHandler)

	dispatcher.RegisterHandler(AiMatchFunc, AiHandler)

	dispatcher.RegisterHandler(MeiNvMatchFunc, MeiNvHandler)

	dispatcher.RegisterHandler(XjjMatchFunc, XjjHandler)

	dispatcher.RegisterHandler(RecalledMatchFunc, RecalledHandler)
	// dispatcher.RegisterHandler(CheckInMatchFunc, CheckInHandler)

	// dispatcher.RegisterHandler(MoYuRiBaoMatchFunc, MoYuRiBaoHandler)

	// dispatcher.RegisterHandler(LeaveMatchFunc, LeaveHandler)

	return dispatcher.AsMessageHandler()
}
