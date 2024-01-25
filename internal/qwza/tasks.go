package qwza

import (
	"github.com/samber/lo"
	"github.com/vikpe/qwza-discord-bot/internal/pkg/task"
	"github.com/vikpe/serverstat"
	"github.com/vikpe/serverstat/qserver"
	"github.com/vikpe/serverstat/qserver/qclient"
)

func NewMonitorTask(serversAddresses []string, onClientsJoined func(server qserver.GenericServer, clients []qclient.Client)) *task.PeriodicalTask {
	statClient := serverstat.NewClient()
	clientIdsPerServer := map[string][]int{}
	isFirstTick := true

	onTick := func() {
		serverInfo := statClient.GetInfoFromMany(serversAddresses)

		for _, server := range serverInfo {
			currentClientIds := lo.Map(server.Clients, func(player qclient.Client, index int) int {
				return player.Id
			})

			newClientIds, _ := lo.Difference(currentClientIds, clientIdsPerServer[server.Address])

			newClients := lo.Filter(server.Clients, func(player qclient.Client, index int) bool {
				return player.Name.ToPlainString() != "[ServeMe]" && lo.Contains(newClientIds, player.Id)
			})

			if len(newClients) > 0 && !isFirstTick {
				onClientsJoined(server, newClients)
			}

			clientIdsPerServer[server.Address] = currentClientIds
		}

		if isFirstTick {
			isFirstTick = false
		}
	}

	return task.NewPeriodicalTask(onTick)
}
