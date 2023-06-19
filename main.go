package main

import (
	"github.com/spf13/viper"
	"log"
	"nokogiriwatir/notifierbot/pkg/client"
	"nokogiriwatir/notifierbot/pkg/eventer"
	"nokogiriwatir/notifierbot/pkg/receiver"
)

func main() {
	if err := setupViper(); err != nil {
		log.Fatal(err)
	}

	host := viper.GetString("HOST")
	token := viper.GetString("TOKEN")

	if host == "" || token == "" {
		log.Printf("There are no token or/and host in .env. Got token: %s, host: %s", token, host)
	}
	c := client.New(viper.GetString("HOST"), viper.GetString("TOKEN"))

	r := receiver.New(*c)

	if er := eventer.ListenEvents(r); er != nil {
		panic(er)
	}
}

func setupViper() error {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if er := viper.ReadInConfig(); er != nil {
		return er
	}

	return nil
}
