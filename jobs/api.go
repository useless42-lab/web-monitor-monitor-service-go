package jobs

import (
	"WebMonitor/models"
	"WebMonitor/monitor"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type CheckApiJob struct {
}

func (checkApiJob CheckApiJob) Run() {
	apiList := models.GetAllActiveApiList()
	for _, item := range apiList {
		monitorRegionArr := strings.Split(item.MonitorRegion, ",")
		monitorRegionResult := InArray(monitorRegionArr, os.Getenv("REGION"))
		if monitorRegionResult {
			apiLogItem := models.GetLatestApiLog(item.Id, os.Getenv("REGION"))
			createdTime := apiLogItem.CreatedAt.Format("2006-01-02 15:04:05")
			createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
			createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(item.Frequency))

			if time.Now().After(createdTimeLocation) {
				var requestData string
				if item.BodyType == 1 {
					requestData = item.BodyRaw
				} else if item.BodyType == 2 {
					requestData = item.BodyJson
				} else if item.BodyType == 3 {
					requestData = item.BodyForm
				}
				status, statusCode, body, proto, elapsed := monitor.MonitorApi(item.Path, item.Method, item.RequestHeaders, item.BodyType, requestData, item.ResponseData, item.BasicUser, item.BasicPassword)
				var checkSuccess int = 0
				if item.ApiMonitorType == 1 {
					// 状态码
					apiHttpStatusCode, _ := strconv.Atoi(item.ApiHttpStatusCode)
					if apiHttpStatusCode == statusCode {
						checkSuccess = 1
					} else {
						checkSuccess = 0
					}
				} else {
					// 返回值
					if body == requestData {
						checkSuccess = 1
					} else {
						checkSuccess = 0
					}
				}

				models.AddApiLog(item.Id, status, statusCode, proto, elapsed, body, checkSuccess, os.Getenv("REGION"))
			} else {
				fmt.Println("未到接口监控时间")
			}
		}
	}
}
