package handler

import (
	"log"
	"net/http"

	"github.com/eatmoreapple/openwechat"
)

func MoYuRiBaoMatchFunc(m *openwechat.Message) bool {
	// TODO cron job
	return false
}

func MoYuRiBaoHandler(ctx *openwechat.MessageContext) {
	url := "https://dayu.qqsuu.cn/moyuribao/apis.php"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	ctx.ReplyImage(resp.Body)
}
