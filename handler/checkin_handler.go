package handler

import (
	"github.com/eatmoreapple/openwechat"
)

func CheckInMatchFunc(m *openwechat.Message) bool {
	return m.Content == "签到"
}
func CheckInHandler(ctx *openwechat.MessageContext) {
	// sender, err := ctx.SenderInGroup()
	// if err != nil {
	// 	log.Println(err)
	// }
	// checkInList, err := query.Checkin.
	// 	Where(query.Checkin.UserID.Eq(sender.ID())).
	// 	Find()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// targetUserID := "2"
	// index := -1
	// for i, checkin := range checkInList {
	// 	if checkin.UserID == targetUserID {
	// 		index = i
	// 		break
	// 	}
	// }
	// if index != -1 {
	// 	fmt.Printf("UserID 为 %s 的数据排第 %d 个\n", targetUserID, index+1)
	// } else {
	// 	fmt.Printf("未找到 UserID 为 %s 的数据\n", targetUserID)
	// }
	// ctx.ReplyText(fmt.Sprintf(`[Fireworks][Fireworks]🏮签到奖励🏮[Fireworks][Fireworks]

	// 🎉昵称：%s
	// 💹排名：第%d名
	// // 💎奖励：1000钻石
	// // 🌈统计：连续4累计4
	// // 🎉时间：10:17
	// ——————————
	// ✨✨♡𝕓𝕖 𝕙𝕒𝕡𝕡𝕪 ➹♡✨✨
	// 有事做   有人爱   有所期待✿
	// ——————————
	// 🎼.•*💌🐾¨*•.¸¸♬🎶•*¨💌
	// 🐾*•.¸¸♬🎼.•*💌🐾
	// 🌸春天比我先拥抱了你
	// @奥德彪 `, sender.NickName, 1))
}
