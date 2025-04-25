package job

import (
	"bot/internal/svc"
	"context"
	"fmt"
)

// type GetTokenIdResp struct {
// 	Ret  int64  `json:"ret"`
// 	Msg  string `json:"msg"`
// 	Data string `json:"data"`
// }
// type GetLoginQrCodeResp struct {
// 	Ret  int64  `json:"ret"`
// 	Msg  string `json:"msg"`
// 	Data Data   `json:"data"`
// }

//	type Data struct {
//		QrData      string `json:"qrData"`
//		QrImgBase64 string `json:"qrImgBase64"`
//		UUID        string `json:"uuid"`
//		AppID       string `json:"appId"`
//	}
type Payload struct {
	AppID      string `json:"appId"`
	ChatroomID string `json:"chatroomId"`
}

type Resp struct {
	Ret  int64  `json:"ret"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	Friends       []string     `json:"friends"`
	Chatrooms     []string     `json:"chatrooms"`
	Ghs           []string     `json:"ghs"`
	MemberList    []MemberList `json:"memberList"`
	ChatroomOwner string       `json:"chatroomOwner"`
	AdminWxid     interface{}  `json:"adminWxid"`
}

type MemberList struct {
	Wxid            string `json:"wxid"`
	NickName        string `json:"nickName"`
	InviterUserName string `json:"inviterUserName"`
	MemberFlag      int64  `json:"memberFlag"`
	DisplayName     string `json:"displayName"`
	BigHeadImgURL   string `json:"bigHeadImgUrl"`
	SmallHeadImgURL string `json:"smallHeadImgUrl"`
}

type Job interface {
	Run()
}

func RegisterJobs(svcCtx *svc.ServiceContext) {
	ctx := context.Background()

	svcCtx.Cron.AddFunc("@every 10s", NewGetChatroomMemberDetailJob(ctx, svcCtx).Run)

	fmt.Printf("Starting job at %s\n", svcCtx.Config.Host)
	svcCtx.Cron.Start()
}
