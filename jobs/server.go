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

type CheckServerJob struct {
}

func (checkServerJob CheckServerJob) Run() {
	serverList := models.GetAllActiveServerList()
	for _, item := range serverList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			serverLogItem := models.GetLatestServerLog(item.Id, os.Getenv("REGION"))
			createdTime := serverLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				// 1 ping 2 服务器状态
				if item.ServerMonitorType == 1 {
					checkSuccess, elapsed := monitor.MonitorServerPing(item.Path)
					var checkSuccessInt int = 0
					if checkSuccess {
						checkSuccessInt = 1
					}
					data := models.ServerLogItem{
						ServerId:          item.Id,
						CpuUser:           "",
						CpuSystem:         "",
						CpuIdle:           "",
						CpuPercent:        "",
						MemoryTotal:       "",
						MemoryAvailable:   "",
						MemoryUsed:        "",
						MemoryUsedPercent: "",
						DiskTotal:         "",
						DiskFree:          "",
						DiskUsed:          "",
						DiskUsedPercent:   "",
						NetSent:           "",
						NetRecv:           "",
						Elapsed:           elapsed,
						CheckSuccess:      checkSuccessInt,
						Region:            os.Getenv("REGION"),
					}
					models.AddServerLog(data)
				}
				if item.ServerMonitorType == 2 {
					endTimeLocation := createdTimeLocation.Add(+time.Second * time.Duration(15))
					checkResult := models.CheckLatestServerLog(item.Id, createdTimeLocation, endTimeLocation)
					if !checkResult {
						data2 := models.ServerLogItem{
							ServerId:          item.Id,
							CpuUser:           "",
							CpuSystem:         "",
							CpuIdle:           "",
							CpuPercent:        "",
							MemoryTotal:       "",
							MemoryAvailable:   "",
							MemoryUsed:        "",
							MemoryUsedPercent: "",
							DiskTotal:         "",
							DiskFree:          "",
							DiskUsed:          "",
							DiskUsedPercent:   "",
							NetSent:           "",
							NetRecv:           "",
							Elapsed:           0,
							CheckSuccess:      0,
							Region:            os.Getenv("REGION"),
						}
						models.AddServerLog(data2)
					}
				}
			} else {
				fmt.Println("未到服务器监控时间")
			}
		}
	}
}
