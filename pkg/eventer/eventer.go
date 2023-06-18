package eventer

import (
	"nokogiriwatir/notifierbot/pkg/cmdHandler"
	"nokogiriwatir/notifierbot/pkg/receiver"
	"time"
)

func ListenEvents(receiver receiver.Receiver) error {
	for {
		updates, er := receiver.Updates(0, 100)

		if er != nil {
			return er
		}

		err := cmdHandler.HandleUpdates(updates)
		if err != nil {
			return err
		}

		time.Sleep(time.Second * 1)
	}
	return nil
}
