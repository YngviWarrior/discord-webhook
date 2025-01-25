package discordService

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

type discord struct{}

type discordInterface interface {
	SendNotification(params *discordstructs.Notification)
}

func NewDiscordWebhook() discordInterface {
	return &discord{}
}

const discordBaseURL = "https://discord.com/api/webhooks"

func (s *discord) SendNotification(params *discordstructs.Notification) {
	client := &http.Client{}

	jsonstr, _ := json.Marshal(params)

	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", discordBaseURL+params.ChannelUrl, bytes.NewBuffer(jsonstr))

	if err != nil {
		log.Println("DSN 01: ", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Discord req exec: ", err)
	}

	defer resp.Body.Close()
}
