package qwza

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"
	"github.com/vikpe/qwza-discord-bot/internal/pkg/discordbot"
	"github.com/vikpe/qwza-discord-bot/internal/qwza/config"
	"github.com/vikpe/serverstat/qserver"
	"github.com/vikpe/serverstat/qserver/convert"
	"github.com/vikpe/serverstat/qserver/qclient"
	"log"
	"time"
)

func New(token string, guildID string, config *config.Config) (*discordbot.Bot, error) {
	bot, err := discordbot.New(token, guildID)

	onPlayersJoined := func(server qserver.GenericServer, clients []qclient.Client) {
		msg := ToPlayersJoinedMessage(server, clients)
		log.Println(msg)
		bot.Say(config.Monitor.ChannelId, msg)
	}

	monitorTask := NewMonitorTask(config.Monitor.Servers, onPlayersJoined)

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

func ToPlayersJoinedMessage(server qserver.GenericServer, clients []qclient.Client) string {
	playerNames := lo.Map(clients, func(client qclient.Client, index int) string {
		return client.Name.ToPlainString()
	})
	playerNamesList := ToNaturalList(playerNames)
	mvdsv := convert.ToMvdsv(server)
	return fmt.Sprintf("%s joined %s - %s on %s (%d/%d, %d specs)", playerNamesList, server.Address, mvdsv.Mode, mvdsv.Settings.Get("map", "unknown"), mvdsv.PlayerSlots.Used, mvdsv.PlayerSlots.Total, mvdsv.SpectatorSlots.Used)
}
