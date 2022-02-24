package models

type OrderItemStruct struct {
	DefaultModel
	UserId         int64   `json:"user_id" gorm:"column:user_id"`
	OrderTime      int64   `json:"order_time" gorm:"column:order_time"`
	OrderAmount    float64 `json:"order_amount" gorm:"column:order_amount"`
	State          string  `json:"state" gorm:"column:state"`
	Payway         string  `json:"payway" gorm:"column:payway"`
	OrderId        string  `json:"order_id" gorm:"column:order_id"`
	UrlKey         string  `json:"url_key" gorm:"column:url_key"`
	TargetPlanId   int     `json:"target_plan_id" gorm:"column:target_plan_id"`
	TargetPlanTime int     `json:"target_plan_time" gorm:"column:target_plan_time"`
	Status         int     `json:"status" gorm:"column:status"`
}

func UseOrder(orderListId int64, userId int64) {
	sqlStr := `
	update order_list set status=0 where id=@orderListId and user_id=@userId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"orderListId": orderListId,
		"userId":      userId,
	})
}

func GetUserOrder(userId int64) OrderItemStruct {
	var result OrderItemStruct
	sqlStr := `select * from order_list where user_id=@userId and status=1 limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}
