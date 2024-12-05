package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	var err error
	// you can change 'bot' and set this by your token
	 TELEGRAM_BOT_TOKEN="7790931634:AAF2MMv4ajrvC26f3dEX41atFWR8ttVwiZA"

	bot, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	//send message by (NumberId_Send_First int )
	/*
 	NumberId_Send_First= (int)
	msg := tgbotapi.NewMessage(NumberId_Send_First, "text")
	for i := 0; i < 10; i++ {

		bot.Send(msg)
	}
	*/
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			//write information of user 
			log.Printf("\nName     : [%s\t%s] \nUsername : [%s] \nMessage  : [%s] \nchat-id  : [%d]", update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)
			//fillter by language
			if update.Message.From.LanguageCode == "en" {

				if update.Message.From.UserName == "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "test")
					for i := 0; i < 40; i++ {

						bot.Send(msg)
					}
				} else {
					var text string
					var usertext string
					usertext = update.Message.Text
					ID := update.Message.Chat.ID
					if strings.HasPrefix(usertext, "/") {
						text = "\n Oh you write a command "
						msg := tgbotapi.NewMessage(ID, "you write : "+usertext+text)
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
						//command 
						switch usertext {
						case "/start":
							start := tgbotapi.NewMessage(ID, "Run now ")
							bot.Send(start)
							break
						case "/profile":
							start := tgbotapi.NewMessage(ID, "Name : "+update.Message.From.FirstName+"\nLast Name : "+update.Message.From.LastName+"\nUserName : "+update.Message.From.UserName)
							bot.Send(start)
							break
						case "/Menu":
							nextButton := "Next"
							firstMenu := "<b>Menu 1</b>\n\nA beautiful menu with a shiny inline button."

							firstMenuMarkup := tgbotapi.NewInlineKeyboardMarkup(
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonData(" تکست", nextButton),

									tgbotapi.NewInlineKeyboardButtonData(" نکست", nextButton),
								),
							)

							msg := tgbotapi.NewMessage(ID, firstMenu)
							msg.ParseMode = tgbotapi.ModeHTML
							msg.ReplyMarkup = firstMenuMarkup
							bot.Send(msg)
							break
						}

					} else {
						// you can change id for owner of bot
						text = "\n this bot not fix you can get screenshot and send for @ekzsooz 🙄"
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "you write : "+update.Message.Text+text)
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)

					}
					//show command
					cmdCfg := tgbotapi.NewSetMyCommands(
						tgbotapi.BotCommand{
							Command:     "/start",
							Description: "starttttt",
						},
						tgbotapi.BotCommand{
							Command:     "/help",
							Description: "helpppp",
						},
						tgbotapi.BotCommand{
							Command:     "/Profile",
							Description: "prooooof",
						},
					)

					bot.Send(cmdCfg)

				}
			} else {
				//another language (persian-فارسی)
				if update.Message.From.UserName == "eksooz" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, " سلام خدا")
					bot.Send(msg)
				} else if update.Message.From.UserName == "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "test")
					for i := 0; i < 7; i++ {

						bot.Send(msg)
					}
				} else {
					var text string
					text = "\nاین بات هنوز کامل نشده اگر اجراش کردین  لطفا اسکرین شات بگیرین و به ایدی زیر ارسال کنید\n @ekzsooz"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "شما تایپ کردید : "+update.Message.Text+text)
					msg.ReplyToMessageID = update.Message.MessageID

					bot.Send(msg)
				}
			}

		}

	}

}
