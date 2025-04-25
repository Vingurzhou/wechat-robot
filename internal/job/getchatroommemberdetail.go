package job

import (
	"bot/internal/dao/model"
	"bot/internal/svc"

	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/clause"
)

type GetChatroomMemberDetailJob struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Run implements Job.
func (j *GetChatroomMemberDetailJob) Run() {
	payload := &Payload{
		AppID: j.svcCtx.Config.GEWE.AppId,
	}
	marshal, err := json.Marshal(payload)
	if err != nil {
		j.Error(err)
		return
	}
	resp, err := j.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/contacts/fetchContactsList", "POST", string(marshal), map[string]string{
		"X-GEWE-TOKEN": j.svcCtx.Config.GEWE.Token,
		"Content-Type": "application/json",
	})
	if err != nil {
		j.Error(err)
		return
	}
	var fetchContactsListResp *Resp
	err = json.Unmarshal(resp, &fetchContactsListResp)
	if err != nil {
		j.Error(err)
		return
	}
	//////////////////////////////////////
	for _, chatroomId := range fetchContactsListResp.Data.Chatrooms {
		payload = &Payload{
			AppID:      j.svcCtx.Config.GEWE.AppId,
			ChatroomID: chatroomId,
		}
		marshal, err = json.Marshal(payload)
		if err != nil {
			j.Error(err)
			return
		}
		resp, err := j.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/group/getChatroomMemberList", "POST", string(marshal), map[string]string{
			"X-GEWE-TOKEN": j.svcCtx.Config.GEWE.Token,
			"Content-Type": "application/json",
		})
		if err != nil {
			j.Error(err)
			return
		}
		var getChatroomMemberListResp *Resp
		err = json.Unmarshal(resp, &getChatroomMemberListResp)
		if err != nil {
			j.Error(err)
			return
		}
		for _, member := range getChatroomMemberListResp.Data.MemberList {
			create := &model.Member{
				Wxid:            member.Wxid,
				NickName:        member.NickName,
				InviterUserName: member.InviterUserName,
				MemberFlag:      int32(member.MemberFlag),
				DisplayName:     member.DisplayName,
				BigHeadImgURL:   member.BigHeadImgURL,
				SmallHeadImgURL: member.SmallHeadImgURL,
				ChatroomID:      chatroomId,
				// IsOwner:         getChatroomMemberListResp.Data.IsOwner,
				// IsAdmin:         getChatroomMemberListResp.Data.IsAdmin,
			}
			err = j.svcCtx.Query.Member.WithContext(j.ctx).
				Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "wxid"}, {Name: "chatroom_id"}},
					UpdateAll: true}).
				Create(create)
			if err != nil {
				j.Error(err)
				return
			}
			// j.Info(create.ID)
		}

	}
}
func NewGetChatroomMemberDetailJob(ctx context.Context, svcCtx *svc.ServiceContext) Job {
	return &GetChatroomMemberDetailJob{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
