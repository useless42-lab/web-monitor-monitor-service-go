package jobs

import "WebMonitor/models"

type DeleteLogJob struct{}

func (deleteLogJob DeleteLogJob) Run() {
	models.DeleteWebLog()
	models.DeleteServerLog()
	models.DeleteApiLog()
	models.DeleteTcpLog()
	models.DeleteDnsLog()
	models.DeleteHeartbeatLog()
	models.DeleteSteamGameServerLog()
	models.DeleteMinecraftServerLog()
}
