package cmd

import (
	"fmt"
	"mdcargas-reporter/internal"
	"mdcargas-reporter/internal/state"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func Run() {
	log := logrus.New()
	cfg, err := internal.GetConfig()
	if err != nil {
		log.Fatal(err.Error()) // no reason to start without proper configuration
	}
	err = state.InitState()
	if err != nil {
		log.Fatal(err.Error())
	}

	bot, err := tg.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Authorized on account %s", bot.Self.UserName)

	IDMap := map[int64]int64{}

	u := tg.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	t := time.NewTicker(time.Second * time.Duration(cfg.Interval))
	for range t.C {
		resp, err := internal.FetchStatus(cfg.Tipo, cfg.Suc, cfg.Numero)
		if err != nil {
			log.Error(err.Error())
			break
		}
		lastState, err := state.ReadState()
		if err != nil {
			log.Error(err.Error())
			break
		}
		if lastState == resp.Estado {
			break
		}

		for update := range updates { // update list of IDs who will receive messages
			if _, ok := IDMap[update.Message.Chat.ID]; !ok {
				IDMap[update.Message.Chat.ID] = update.Message.Chat.ID
			}
		}

		for _, v := range IDMap {
			bot.Send(tg.NewMessage(v, fmt.Sprintf("New state: %s", resp.Estado)))
		}

		// push estado por telegram
		err = state.WriteState(resp.Estado)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
