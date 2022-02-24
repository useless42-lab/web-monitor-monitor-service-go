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

type CheckTcpJob struct {
}

func (checkTcpJob CheckTcpJob) Run() {
	tcpList := models.GetAllActiveTcpList()
	for _, item := range tcpList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			tcpLogItem := models.GetLatestTcpLog(item.Id, os.Getenv("REGION"))
			createdTime := tcpLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				checkSuccess, elapsed := monitor.MonitorTcpPing(item.Path)
				var checkSuccessInt int = 0
				if checkSuccess {
					checkSuccessInt = 1
				}
				models.AddTcpLog(item.Id, elapsed, checkSuccessInt, os.Getenv("REGION"))
			} else {
				fmt.Println("未到TCP监控时间")
			}
		}
	}
}
