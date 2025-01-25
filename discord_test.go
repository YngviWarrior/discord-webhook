package discordService_test

import (
	"os"
	"testing"

	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
)

var ds discordService.DiscordInterface

func TestMain(m *testing.M) {
	ds = discordService.NewDiscordWebhook()

	code := m.Run()
	os.Exit(code)
}

func TestNewDiscordWebhook(t *testing.T) {
	ds.SendNotification(&discordstructs.Notification{})
}
