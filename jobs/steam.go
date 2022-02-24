package jobs

import (
	"WebMonitor/models"
	"WebMonitor/monitor"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type CheckSteamGameServerJob struct {
}

func (checkSteamGameServerJob CheckSteamGameServerJob) Run() {
	steamGameServerList := models.GetAllActiveSteamGameServerList()
	for _, item := range steamGameServerList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			steamGameServerLogItem := models.GetLatestSteamGameServerLog(item.Id, os.Getenv("REGION"))
			createdTime := steamGameServerLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				checkSuccess, name, playersMax, playersOnline, elapsed := monitor.MonitorSteam(item.SteamApiKey, item.Path)
				var checkSuccessInt int = 0
				if checkSuccess {
					checkSuccessInt = 1
				}
				models.AddSteamGameServerLog(item.Id, name, playersMax, playersOnline, int(elapsed), checkSuccessInt, os.Getenv("REGION"))
			} else {
				fmt.Println("未到STEAMGAMESERVER监控时间")
			}
		}
	}
}
