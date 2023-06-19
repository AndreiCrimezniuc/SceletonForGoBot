package client

import "nokogiriwatir/notifierbot/pkg/entities"

type Update struct {
	ID      int64            `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string        `json:"text"`
	User entities.User `json:"from"`
	Chat entities.Chat `json:"chat"`
	ID   int64         `json:"message_id"`
}
