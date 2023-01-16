package application

import (
	"dgb/meter.readings.bot/internal/configuration"
	"dgb/meter.readings.bot/internal/database"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func StartTelegramBot(conf configuration.Configuration, repo database.Repository) {

	bot, err := tgbotapi.NewBotAPI(conf.TELEGRAM_BOT_TOKEN)

	if err != nil {
		log.Panic(err)
	}

	startTelegramBot(conf, bot, repo)
}

func startTelegramBot(conf configuration.Configuration, bot *tgbotapi.BotAPI, repo database.Repository) {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 6

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil || !update.Message.IsCommand() || !checkFrom(update.Message.From, conf.TELEGRAM_VALID_USERS) {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start", "help":
			msg.Text = "Hello, please use the commands /add n:<reading> d:<reading>, /latest and /help."
		case "add":
			t := strings.Trim(strings.ReplaceAll(update.Message.Text, "/add", ""), " ")
			valid, readingErr := addReading(repo, t, update.Message.From.UserName)

			if !valid {
				msg.Text = fmt.Sprintf("Sorry, I couldn't add your readings %s", t)
				log.Printf("[%s] %s", update.Message.From.UserName, readingErr)
			}

			msg.Text = fmt.Sprintf("Ok, I've added your readings %s", t)
		case "latest":
			readings := getReadings(repo)
			msg.Text = fmt.Sprintf("The latest readings are:\r\n%s", readings)
		// case "open":
		// 	msg.ReplyMarkup = numericKeyboard
		// 	msg.Text = update.Message.Text
		default:
			msg.Text = "I'm afraid I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func checkFrom(user *tgbotapi.User, validUsers []int64) bool {
	for _, s := range validUsers {
		if user.ID == s {
			return true
		}
	}

	return false
}

func getReadings(repo database.Repository) string {
	res := repo.GetLatest()
	str := ""

	for _, s := range res {
		str = str + fmt.Sprintf("%v:%v\r\n", s["rate"], s["reading"])
	}

	return str
}

func addReading(repo database.Repository, readings string, user string) (bool, string) {

	split := strings.Split(readings, " ")
	d, dayConvErr := strconv.Atoi(strings.Split(split[0], ":")[1])
	n, nightConvErr := strconv.Atoi(strings.Split(split[1], ":")[1])

	if dayConvErr != nil {
		return false, dayConvErr.Error()
	}

	if nightConvErr != nil {
		return false, nightConvErr.Error()
	}

	date := time.Now().Format(time.RFC3339Nano)

	_, err := repo.Insert(bson.M{
		"_id":         uuid.New().String(),
		"rate":        "Night",
		"note":        "Added via Telegram bot",
		"readingdate": date,
		"reading":     n,
		"userName":    user,
	})

	if err != nil {
		return false, err.Error()
	}

	_, err = repo.Insert(bson.M{
		"_id":         uuid.New().String(),
		"rate":        "Day",
		"note":        "Added via Telegram bot",
		"readingdate": date,
		"reading":     d,
		"userName":    user,
	})

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

// var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("/add", "/add"),
// 		tgbotapi.NewInlineKeyboardButtonData("/help", "/help"),
// 	),
// )
