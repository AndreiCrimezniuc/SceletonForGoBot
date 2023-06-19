package sender

import (
	"context"
	"net/url"
	"nokogiriwatir/notifierbot/pkg/client"
	"strconv"
)

type Sender struct {
	Client *client.Client
}

const SendMessageMethod = "sendMessage"

func (s *Sender) SendMessage(ctx context.Context, chatID int64, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.FormatInt(chatID, 10))
	q.Add("text", text)

	_, err := s.Client.DoRequest(ctx, SendMessageMethod, q)
	if err != nil {
		return err
	}

	return nil
}
