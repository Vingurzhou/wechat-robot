package handler

import (
	"log"
	"net/http"

	"github.com/eatmoreapple/openwechat"
)

func XjjMatchFunc(m *openwechat.Message) bool {
	return m.Content == "康康美女视频"
}
func XjjHandler(ctx *openwechat.MessageContext) {
	url := "https://api.lolimi.cn/API/xjj/xjj.php"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	ctx.ReplyVideo(resp.Body)
}
