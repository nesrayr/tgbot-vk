package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"tgbot/database"
	"tgbot/handlers"
	"tgbot/models"
	"time"
)

func main() {
	database.ConnectDB()
	bot, err := tgbotapi.NewBotAPIWithClient(os.Getenv("BOT_TOKEN"), &http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	})
	if err != nil {
		fmt.Println("not correct token")
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	var infos []models.Info
	for update := range updates {
		database.DB.Db.Where("expiration < ?", time.Now()).Delete(&infos)
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.Chat.UserName)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "set":
			handlers.HandleSetCommand(bot, update)
		case "get":
			handlers.HandleGetCommand(bot, update)
		case "del":
			handlers.HandleDelCommand(bot, update)
		}
	}
}
