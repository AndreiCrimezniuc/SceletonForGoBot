package main

import (
	"flag"
	"nokogiriwatir/notifierbot/pkg/client"
	"nokogiriwatir/notifierbot/pkg/eventer"
	"nokogiriwatir/notifierbot/pkg/receiver"
)

const thBotHost = "api.telegram.org"

func main() {
	token := flag.String(
		"token",
		"",
		"token to access tg bot",
	)

	flag.Parse()

	c := client.New(thBotHost, *token)

	r := receiver.New(*c)

	if er := eventer.ListenEvents(*r); er != nil {
		panic(er)
	}
}
