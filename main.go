package main

import (
	"log"
	"os"

	widget "github.com/ketchupsalt/slack-widget"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func main() {
	apiKey := os.Getenv("SLACK_XOXB")
	if apiKey == "" {
		log.Fatalf("set SLACK_XOXB to run")
	}

	url := os.Getenv("LISTEN_URL")
	if url == "" {
		url = "http://localhost:3000/events-endpoint"
	}


	bot, err := widget.New(apiKey, url)
	if !widget.OK(err) {
		return
	}

	log.Printf("[%s] listening on %s", bot.User, url)

	for iev := range bot.Events {
		switch ev := iev.Data.(type) {
		case *slackevents.AppMentionEvent:
			if ev.User != bot.User {
				bot.API.PostMessage(ev.Channel, slack.MsgOptionText("*#leetcode HOF!*\n* @grepcats", false))
			}
			log.Printf("[%s] <%s> %s", bot.GetChannelName(ev.Channel), bot.GetUserName(ev.User), ev.Text)
		}
	}
}