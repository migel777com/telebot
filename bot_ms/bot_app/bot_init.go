package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	_ "os"
)

func main(){
	//load global variables
	_ = godotenv.Load("globals.env")
	bot, err := tgbotapi.NewBotAPI("1627956704:AAH7TeX80Xy_TZGe0Z4aOtVwfXC_3XsSQdM")
	if err != nil{
		log.Println("Error when creating a bot instance occured", err)
	}
	log.Println("TG bot is running")
	Bot_handler_init(bot)
}