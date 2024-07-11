package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eatmoreapple/openwechat"
)

type MeiNvResponse struct {
	Code int64  `json:"code"`
	Text string `json:"text"`
	Data struct {
		Image string `json:"image"`
	} `json:"data"`
}

func MeiNvMatchFunc(m *openwechat.Message) bool {
	return m.Content == "康康美女图片"
}
func MeiNvHandler(ctx *openwechat.MessageContext) {
	url := "https://api.lolimi.cn/API/meinv/api.php?type=json"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var response MeiNvResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return
	}

	resp2, err := http.Get(response.Data.Image)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp2.Body.Close()
	ctx.ReplyImage(resp2.Body)
}
