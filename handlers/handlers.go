package handlers

import (
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"tgbot/database"
	"tgbot/models"
	"time"
)

func HandleSetCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	args := update.Message.CommandArguments()
	if args == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный формат. Должно быть: /set <service> <login> <password>")
		bot.Send(msg)
		return
	}
	fields := strings.Fields(args)
	if len(fields) != 3 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный формат. Должно быть: /set <service> <login> <password>")
		bot.Send(msg)
		return
	}
	serviceName := fields[0]
	login := fields[1]
	password := hex.EncodeToString([]byte(fields[2]))
	info := models.Info{ServiceName: serviceName, Login: login, Password: password, UserName: update.Message.From.UserName,
		Expiration: time.Now().Add(24 * time.Hour * 3)} // Данные будут храниться 3 суток
	if err := database.DB.Db.Create(&info).Error; err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Логин и пароль для сервиса %s успешно сохранены", serviceName))
	bot.Send(msg)
}

func HandleGetCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	serviceName := update.Message.CommandArguments()
	var info models.Info
	if err := database.DB.Db.Where("service_name=? AND user_name=?", serviceName, update.Message.From.UserName).
		First(&info).Error; err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
		return
	}
	password, err := hex.DecodeString(info.Password)
	if err != nil {
		fmt.Println(err)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Логин: %s\nПароль: %s", info.Login, password))
	bot.Send(msg)
}

func HandleDelCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	serviceName := update.Message.CommandArguments()
	var info models.Info
	if err := database.DB.Db.Where("service_name=? AND user_name=?", serviceName, update.Message.From.UserName).
		Delete(&info).Error; err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Логин и пароль сервиса %s успешно удалены", serviceName))
	bot.Send(msg)
}
