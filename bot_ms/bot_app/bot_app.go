package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	_ "os"
	"strconv"
)


func connectToMainServer(pattern string, chat_id string) (response string){

	fmt.Println("ClientConnected")
	_ = godotenv.Load("globals.env")
	/*port := os.Getenv("MAIN_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")*/
	port := "59751"
	address := "0.0.0.0"
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := NewTelegramBotServiceClient(conn)
	ctx := context.Background()
	req := &CommandPackRequest{Command: pattern, ChatId: chat_id}

	res, err := c.CommandPack(ctx, req)
	if err!=nil{
		log.Fatalln("Error when getting response from main_server: ",err)
	}
	response = res.Response
	return response

}

func Bot_handler_init(bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln("Error when listening to updates: ",err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		response_string := connectToMainServer(update.Message.Text, strconv.FormatInt(update.Message.Chat.ID, 10))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response_string)
		_, err = bot.Send(msg)
		if err!=nil{
			log.Println("Error when sending via tg bot: ",err)
		}
	}
}






