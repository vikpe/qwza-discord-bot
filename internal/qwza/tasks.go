package qwza

import (
	"github.com/samber/lo"
	"github.com/vikpe/qwza-discord-bot/internal/pkg/task"
	"github.com/vikpe/serverstat"
	"github.com/vikpe/serverstat/qserver"
	"github.com/vikpe/serverstat/qserver/qclient"
)

func NewMonitorTask(serversAddresses []string, onPlayersJoined func(server qserver.GenericServer, clients []qclient.Client)) *task.PeriodicalTask {
	statClient := serverstat.NewClient()
	playerIdsPerServer := map[string][]int{}
	isFirstTick := true

	onTick := func() {
		serverInfo := statClient.GetInfoFromMany(serversAddresses)

		for _, server := range serverInfo {
			currentPlayerIds := lo.Map(server.Clients, func(player qclient.Client, index int) int {
				return player.Id
			})

			newPlayerIds, _ := lo.Difference(currentPlayerIds, playerIdsPerServer[server.Address])

			newPlayers := lo.Filter(server.Players(), func(player qclient.Client, index int) bool {
				return lo.Contains(newPlayerIds, player.Id)
			})

			if len(newPlayers) > 0 && !isFirstTick {
				onPlayersJoined(server, newPlayers)
			}

			playerIdsPerServer[server.Address] = currentPlayerIds
		}

		if isFirstTick {
			isFirstTick = false
		}
	}

	return task.NewPeriodicalTask(onTick)
}
