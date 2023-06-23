package cmdHandler

import (
	"github.com/spf13/viper"
	"log"
	"nokogiriwatir/notifierbot/pkg/client"
	"nokogiriwatir/notifierbot/pkg/entities"
	"nokogiriwatir/notifierbot/pkg/receiver"
	SenderObj "nokogiriwatir/notifierbot/pkg/sender"
)

const (
	helloCMD   = "/hello"
	goodbyeCMD = "/goodbye"
	startCMD   = "/start"
	helpCMD    = "/help"
)

type sender interface {
	Send()
}

type fetcher interface {
	fetch()
}

func HandleUpdates(updates []client.Update) error {
	host := viper.GetString("HOST")
	token := viper.GetString("TOKEN")

	for _, val := range updates {

		if val.ID == receiver.GetLastMessageID() {
			continue
		}

		handleCMD(
			val.Message.Text,
			&SenderObj.Sender{
				Client: client.New(host, token),
			},
			&val.Message.User,
			&val.Message.Chat)
		receiver.SetLastMessageID(val.ID)
	}
	return nil
}

func handleCMD(cmd string, sender *SenderObj.Sender, user *entities.User, chat *entities.Chat) {
	switch cmd {
	case helloCMD:
		handleHello()
	case goodbyeCMD:
		handleGoodBye()
	case startCMD, helpCMD:
		handleStart(sender, user, chat)
	}

	//log.Print("there is no command like that")
}

//func handleHelp() {
//	handleStart()
//}

func handleGoodBye() {
	log.Print("Goodbye too, man!")
}

func handleHello() {
	log.Print("Hello too, man!")
}
