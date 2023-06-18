package sender

import (
	"context"
	"net/url"
	"nokogiriwatir/notifierbot/pkg/client"
	"strconv"
)

type Sender struct {
	client *client.Client
}

const SendMessageMethod = "sendMessage"

func (s *Sender) SendMessage(ctx context.Context, chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := s.client.DoRequest(ctx, SendMessageMethod, q)
	if err != nil {
		return err
	}

	return nil
}
