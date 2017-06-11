package notifier

import (
	"fmt"
  "os"
	"github.com/nlopes/slack"
)

type Slack struct {}

var TOKEN = os.Getenv("AMAZON_ORDER_SLACK_TOKEN")
var CHANNEL_ID = os.Getenv("AMAZON_ORDER_CHANNEL_ID")

func (c Slack)PostMessage(message string) {
	api := slack.New(TOKEN)
	params := slack.PostMessageParameters{}
	channelID, timestamp, err := api.PostMessage(CHANNEL_ID, message, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
