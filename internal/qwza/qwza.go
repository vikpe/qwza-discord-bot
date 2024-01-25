package qwza

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/vikpe/qwza-discord-bot/internal/pkg/discordbot"
	"github.com/vikpe/qwza-discord-bot/internal/qwza/config"
	"github.com/vikpe/qwza-discord-bot/internal/qwza/tasks/monitor"
	"github.com/vikpe/serverstat/qserver"
	"github.com/vikpe/serverstat/qserver/qclient"
	"log"
	"time"
)

func New(token string, guildID string, config *config.Config) (*discordbot.Bot, error) {
	bot, err := discordbot.New(token, guildID)

	onPlayersJoined := func(server qserver.GenericServer, clients []qclient.Client) {
		log.Println(fmt.Sprintf("Players joined %s: %v", server.Address, clients))
	}

	monitorTask := monitor.New(config.Monitor.Servers, onPlayersJoined)

	bot.OnReady = func(s *discordgo.Session) {
		log.Println(fmt.Sprintf("%s is ready", s.State.User.Username))
		monitorTask.Start(time.Duration(config.Monitor.Interval) * time.Second)
	}

	bot.OnStop = func() {
		log.Println("Shutting down.")
		monitorTask.Stop()
	}

	return bot, err
}
