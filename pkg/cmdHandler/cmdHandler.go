package cmdHandler

import (
	"errors"
	"log"
	"nokogiriwatir/notifierbot/pkg/client"
)

const (
	helloCMD   = "/hello"
	goodbyeCMD = "/goodbye"
	startCMD   = "/start"
)

type sender interface {
	Send()
}

type fetcher interface {
	fetch()
}

func HandleUpdates(updates []client.Update) error {
	for _, val := range updates {
		handleCMD(val.Message.Text)
	}
	return nil
}

func handleCMD(cmd string) error {
	switch cmd {
	case helloCMD:
		handleHello()
	case goodbyeCMD:
		handleGoodBye()
	case startCMD:
		start()
	}

	return errors.New("there is no command like that")
}

func handleGoodBye() {
	log.Print("Goodbye too, man!")
}

func start() {

}

func handleHello() {
	log.Print("Hello too, man!")
}
