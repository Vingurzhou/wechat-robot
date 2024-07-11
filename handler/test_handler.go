package handler

import (
	"log"

	"github.com/eatmoreapple/openwechat"
)

func TestMatchFunc(m *openwechat.Message) bool {
	sender, _ := m.Sender()
	senderInGroup, _ := m.SenderInGroup()
	log.Println(m.MsgId, m.MsgType, sender, senderInGroup, m.Content)
	return sender.NickName == "test"
}
func TestHandler(ctx *openwechat.MessageContext) {

	// groups, err := ctx.Owner().Groups()
	// if err != nil {
	// 	log.Println(err)
	// }
	// for _, group := range groups {
	// 	err := query.Group.Clauses(clause.OnConflict{UpdateAll: true}).Create(&model.Group{
	// 		ID:          group.ID(),
	// 		Username:    group.UserName,
	// 		Nickname:    group.NickName,
	// 		DisplayName: group.DisplayName,
	// 	})
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	members, err := group.Members()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	for _, member := range members {
	// 		err := member.Detail()
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		log.Println(member.ID())
	// 		err = query.User.Clauses(clause.OnConflict{UpdateAll: true}).Create(&model.User{
	// 			ID:          model.GenerateSnowflakeID(),
	// 			GroupID:     group.ID(),
	// 			Username:    member.UserName,
	// 			Nickname:    member.NickName,
	// 			DisplayName: member.DisplayName,
	// 		})
	// 		if err != nil && err.Error() != "UNIQUE constraint failed: user.group_id, user.nickname" {
	// 			log.Println(err)
	// 		}
	// 	}
	// }
}
