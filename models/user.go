package models

import "time"

type User struct {
	DefaultModel
	// Username  string    `json:"username" gorm:"column:username"`
	// Email     string    `json:"email" gorm:"column:email" gorm:"unique"`
	// Password  string    `json:"-" gorm:"column:password"`
	// Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Username  string    `json:"username" gorm:"column:username"`
	PlanId    int       `json:'plan_id gorm:"column:plan_id"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func GetAllUserList() []User {
	var result []User
	sqlStr := `select * from user`
	DB.Raw(sqlStr).Scan(&result)
	return result
}

func ResetUserPlan(userId int64) {
	sqlStr := `update user set plan_id=1  where id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		// "expiredAt": time.Now().AddDate(0, 1, 0),
		"userId": userId,
	})
}
func UpdateUserPlan(userId int64, planId int, expiredAt time.Time) {
	sqlStr := `update user set plan_id=@planId , expired_at=@expiredAt where id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"userId":    userId,
		"expiredAt": expiredAt,
		"planId":    planId,
	})
}
