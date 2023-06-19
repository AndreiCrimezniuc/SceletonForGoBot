package eventer

import (
	"nokogiriwatir/notifierbot/pkg/cmdHandler"
	recieverPkg "nokogiriwatir/notifierbot/pkg/receiver"
	"time"
)

func ListenEvents(receiver *recieverPkg.Receiver) error {
	for {
		updates, er := receiver.Updates(-1, 1)

		if er != nil {
			return er
		}

		if updates != nil {
			err := cmdHandler.HandleUpdates(updates)
			if err != nil {
				return err
			}
		}

		time.Sleep(time.Second * 1)
	}
	return nil
}
