package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"google.golang.org/grpc"
	"homework-2/config"
	"homework-2/localusers"
	"homework-2/pkg/api"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	client := api.NewAwesomeBotIIIClient(conn)
	ctx := context.Background()

	rawСfg, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Println(err)
		rawСfg, err = ioutil.ReadFile("/app/config/config.yaml")
		if err != nil {
			log.Println(err)
		}
	}
	cfg, err := config.ParseConfig(rawСfg)
	if err != nil {
		log.Println(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.ApiKeys.Telegram)
	if err != nil {
		log.Panic(err)
	}

	users := localusers.ReadUsers()

	go func() {
		for {
			for i := range users {
				respSrc, err := client.GetSrcsByChat(ctx, &api.ChData{ChatID: users[i]})
				if err != nil {
					log.Println(err)
				}
				for j := range respSrc.Sources {
					respRSS, err := client.GetRSSBySource(ctx, &api.ChSrcData{ChatID: users[i], Source: respSrc.Sources[j]})
					if err != nil {
						log.Println(err)
						continue
					}
					for k := range respRSS.News {
						_, err := bot.Send(tgbotapi.NewMessage(users[i], respRSS.News[k]))
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
			time.Sleep(time.Minute)
		}
	}()

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		log.Println(update.Message.Command())
		log.Println(update.Message.CommandArguments())

		if update.Message.Command() != "" {
			switch update.Message.Command() {
			case "help":
				_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("/start - добавление в базу;\n"+
					"/add linkSrc - добавление источника;\n"+
					"/del linkSrc - удаление источника;\n"+
					"/sources - показать список всех источников;")))
				if err != nil {
					log.Println(err)
				}

			case "start":
				resp, err := client.CreateUser(ctx, &api.ChData{ChatID: update.Message.Chat.ID})
				if err != nil {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать, снова."))
					if err != nil {
						log.Println(err)
					}
				} else {
					users = append(users, update.Message.Chat.ID)
					localusers.WriteUsers(users)
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать путник."))
					if err != nil {
						log.Println(err)
					}
				}
				fmt.Printf("Respond: <%v>\n", resp.String())

			case "add":
				resp, err := client.CreateSource(ctx, &api.ChSrcData{ChatID: update.Message.Chat.ID, Source: update.Message.CommandArguments()})
				if err != nil {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
					if err != nil {
						log.Println(err)
					}
				} else {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Источник добавлен."))
					if err != nil {
						log.Println(err)
					}
				}
				fmt.Printf("Respond: <%v>\n", resp.String())

			case "sources":
				resp, err := client.GetSrcsByChat(ctx, &api.ChData{ChatID: update.Message.Chat.ID})
				if err != nil {
					_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не повезло :("))
					if err != nil {
						log.Println(err)
					}
				} else {
					if resp == nil {
						_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ничего не найдено, не повезло."))
						if err != nil {
							log.Println(err)
						}
					}
					for i := range resp.Sources {
						_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, resp.Sources[i]))
						if err != nil {
							log.Println(err)
						}
					}
				}
				fmt.Printf("Respond: <%v>\n", resp.String())

			case "del":
				resp, err := client.DeleteSource(ctx, &api.ChSrcData{ChatID: update.Message.Chat.ID, Source: update.Message.CommandArguments()})
				if err != nil {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
					if err != nil {
						log.Println(err)
					}
				} else {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Источник удален."))
					if err != nil {
						log.Println(err)
					}
				}
				fmt.Printf("Respond: <%v>\n", resp.String())

			default:
				_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Доступные комманды: /start, /add, /sources, /deletesource"))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
