package jobs

import (
	"WebMonitor/models"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type CheckHeartbeatJob struct{}

func (checkHeartbeatJob CheckHeartbeatJob) Run() {
	heartbeatList := models.GetAllActiveHeartbeatList()
	for _, item := range heartbeatList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			heartbeatLogItem := models.GetLatestHeartbeatLog(item.Id)
			createdTime := heartbeatLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))
			endTimeLocation := createdTimeLocation.Add(+time.Second * time.Duration(15))
			checkResult := models.CheckLatestHeartbeatLog(item.Id, createdTimeLocation, endTimeLocation)
			if !checkResult {
				models.AddHeartbeatLog(item.Id, "", 0)
			}
		}
	}
}
