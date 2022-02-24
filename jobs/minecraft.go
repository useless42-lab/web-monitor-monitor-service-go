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

type CheckMinecraftServerJob struct {
}

func (checkMinecraftServerJob CheckMinecraftServerJob) Run() {
	minecraftServerList := models.GetAllActiveMinecraftServerList()
	for _, item := range minecraftServerList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			minecraftServerLogItem := models.GetLatestMinecraftServerLog(item.Id, os.Getenv("REGION"))
			createdTime := minecraftServerLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				itemPathArr := strings.Split(item.Path, ":")
				if len(itemPathArr) > 1 {
					checkSuccess, playersMax, playersOnline, versionName, elapsed := monitor.MonitorMinecraft(item.PlatformVersion, itemPathArr[0])
					var checkSuccessInt int = 0
					if checkSuccess {
						checkSuccessInt = 1
					}
					models.AddMinecraftServerLog(item.Id, playersMax, playersOnline, versionName, int(elapsed), os.Getenv("REGION"), checkSuccessInt)
				}
			} else {
				fmt.Println("未到MINECRAFTSERVER监控时间")
			}
		}
	}
}
