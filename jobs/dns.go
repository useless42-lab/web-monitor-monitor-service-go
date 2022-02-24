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

type CheckDnsJob struct {
}

func (checkDnsJob CheckDnsJob) Run() {
	dnsList := models.GetAllActiveDnsList()
	for _, item := range dnsList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			dnsLogItem := models.GetLatestDnsLog(item.Id, os.Getenv("REGION"))
			createdTime := dnsLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				checkSuccess, responseData, elapsed := monitor.MonitorDns(item.Path, item.DnsType, item.DnsServer)
				var checkSuccessInt int = 0
				if checkSuccess {
					checkSuccessInt = 1
				}
				models.AddDnsLog(item.Id, item.DnsType, elapsed, responseData, checkSuccessInt, os.Getenv("REGION"))
			} else {
				fmt.Println("未到DNS监控时间")
			}
		}
	}
}
