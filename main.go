package main

import (
	"bybit/bybit/bybit"
	"bybit/bybit/get"
	"bybit/bybit/post"
	"bybit/bybit/print"
	"bybit/bybit/telegram"
	"bybit/env"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var api env.Env
	var trade bybit.Trades
	// var trade map[string]bybit.Trade

	err := env.LoadEnv(&api)
	if err != nil {
		log.Fatalf("Error cannot Read file .env")
	}
	log.Printf("Get api Ok")
	botapi, err := tgbotapi.NewBotAPI(api.Api_telegram)
	if err != nil {
		log.Panic(err)
	}

	botapi.Debug = true

	log.Printf("Authorized on account %s", botapi.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botapi.GetUpdatesChan(u)

	for update := range updates {
		if update.ChannelPost != nil {
			msg := update.ChannelPost.Text
			dataBybite, err := telegram.ParseMsg(msg)
			if err == nil && dataBybite.Trade {
				price := get.GetPrice(dataBybite.Currency)
				log.Println(print.PrettyPrint(price))
				if price.RetCode == 0 {
					trade.Add(api, dataBybite, price)
					err := post.PostOrder(dataBybite.Currency, api, &trade)
					if err != nil {
						log.Println(err)
					}

					log.Println(print.PrettyPrint(trade))
					log.Println(print.PrettyPrint(dataBybite))
					trade.Print()
				}
			} else if err == nil && dataBybite.Cancel {
				log.Printf("Error Parsing")
			}
		}
	}
}
