package jobs

import (
	"WebMonitor/models"
	"strconv"
	"time"
)

type CheckPlanJob struct{}

func (checkPlanJob CheckPlanJob) Run() {
	userList := models.GetAllUserList()
	for _, item := range userList {
		userOrder := models.GetUserOrder(item.DefaultModel.ID)
		if item.PlanId != 1 && time.Now().After(item.ExpiredAt) {
			if userOrder.OrderId == "" {
				models.ResetUserPlan(item.DefaultModel.ID)
				result := models.GetAllDeviceListOwnerByUserId(item.ID)
				for _, item1 := range result {
					deviceId, _ := strconv.ParseInt(item1.Id, 10, 64)
					models.ResetDeviceOwnerByUser(deviceId, item1.DeviceType)
				}
			} else {
				models.UpdateUserPlan(item.DefaultModel.ID, userOrder.TargetPlanId, time.Now().AddDate(0, userOrder.TargetPlanTime, 0))
				models.UseOrder(userOrder.DefaultModel.ID, item.DefaultModel.ID)
			}
		}
	}
}
