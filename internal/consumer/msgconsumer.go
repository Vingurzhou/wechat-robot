package consumer

import (
	"bot/internal/svc"

	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/genai"
)

type MsgConsumer struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgConsumer(context context.Context, serverCtx *svc.ServiceContext) Consumer {
	return &MsgConsumer{
		Logger: logx.WithContext(context),
		ctx:    context,
		svcCtx: serverCtx,
	}
}

type TextMsg struct {
	AppID   string `json:"appId"`
	ToWxid  string `json:"toWxid"`
	Ats     string `json:"ats"`
	Content string `json:"content"`
}

// Consume implements Consumer.
func (c *MsgConsumer) Consume() error {
	for {
		select {
		case msg := <-c.svcCtx.MsgChannel:
			if msg != nil {
				// curl --location --request POST 'http://127.0.0.1:2531/v2/api/message/postText' \
				// --header 'X-GEWE-TOKEN: 447d53fdf6354f1791531049fdba865c' \
				// --header 'Content-Type: application/json' \
				// --data-raw '{
				//     "appId": "wx_6UI6kq1iECVFmp5pqQBp4",
				//     "toWxid": "wxid_vxgg4h0scfmi12",
				//     "ats": "wxid_vxgg4h0scfmi12",
				//     "content": "@猿猴 我在测试艾特内容"
				// }'
				result, err := c.svcCtx.GenAICli.Models.GenerateContent(
					c.ctx,
					"gemini-2.0-flash",
					genai.Text("Explain how AI works in a few words"),
					nil,
				)
				if err != nil {
					c.Error(err)
					continue
				}
				textMsg := TextMsg{
					AppID:   c.svcCtx.Config.GEWE.AppId,
					ToWxid:  msg.FromUserName,
					Ats:     msg.FromUserName,
					Content: result.Text(),
				}
				marshal, err := json.Marshal(textMsg)
				if err != nil {
					c.Error(err)
					continue
				}
				headers := map[string]string{
					"X-GEWE-TOKEN": "447d53fdf6354f1791531049fdba865c",
					"Content-Type": "application/json",
				}
				resp, err := c.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/message/postText", "POST", string(marshal), headers)
				if err != nil {
					c.Error(err)
					continue
				}
				c.Info(string(resp))
			}
		case <-c.ctx.Done():
			return nil
		}
	}
}
