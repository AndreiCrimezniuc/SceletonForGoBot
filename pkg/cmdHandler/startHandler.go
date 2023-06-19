package cmdHandler

import (
	"context"
	"fmt"
	"log"
	"nokogiriwatir/notifierbot/pkg/entities"
	SenderObj "nokogiriwatir/notifierbot/pkg/sender"
	"strings"
)

const greetingMessage = `
	Привет, %s. Доступны следующие команды
	%s
`

func CommandList() []string {
	return []string{
		helloCMD,
		startCMD,
		helloCMD,
	}
}

func handleStart(sender *SenderObj.Sender, user *entities.User, chat *entities.Chat) error {
	log.Printf("Command to handle: %s", startCMD)

	return sender.SendMessage(context.Background(), chat.ID, fmt.Sprintf(greetingMessage, user.FirstName,
		strings.Join(CommandList(), "\n")))
}
