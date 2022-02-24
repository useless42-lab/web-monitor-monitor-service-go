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

type CheckWebJob struct {
}

func (checkWebJob CheckWebJob) Run() {
	webList := models.GetAllActiveWebList()
	for _, item := range webList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			webLogItem := models.GetLatestWebLog(item.Id, os.Getenv("REGION"))
			createdTime := webLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				status, statusCode, body, proto, elapsed := monitor.MonitorWeb(item.Path, item.BasicUser, item.BasicPassword)
				var checkSuccess int
				if item.WebMonitorType == 1 {
					// http状态码监控
					if item.WebHttpStatusCode == statusCode {
						checkSuccess = 1
					} else {
						checkSuccess = 0
					}
				}
				if item.WebMonitorType == 2 {
					// 关键词监控
					if strings.Contains(body, item.WebHttpRegexpText) {
						checkSuccess = 1
					} else {
						checkSuccess = 0
					}
				}
				models.AddWebLog(item.Id, status, statusCode, proto, elapsed, body, checkSuccess, os.Getenv("REGION"))
			} else {
				fmt.Println("未到网站监控时间")
			}
		}
	}
}
