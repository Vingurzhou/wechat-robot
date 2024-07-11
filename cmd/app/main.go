package main

import (
	"log"
	"os"

	"github.com/Vingurzhou/wechat-robot/db"
	"github.com/Vingurzhou/wechat-robot/handler"
	"github.com/Vingurzhou/wechat-robot/query"

	"github.com/eatmoreapple/openwechat"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[wechat-bot] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	query.SetDefault(db.NewDatabase().GormDB)
}
func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	bot.MessageHandler = handler.NewMessageHandler()

	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// bot.Login()

	bot.Block()

}
