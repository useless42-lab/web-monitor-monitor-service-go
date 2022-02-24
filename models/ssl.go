package models

import "time"

type RSSLItem struct {
	WebId     string    `json:"web_id" gorm:"column:web_id"`
	StartTime LocalTime `json:"start_time" gorm:"column:start_time"`
	EndTime   LocalTime `json:"end_time" gorm:"column:end_time"`
	Subject   string    `json:"subject" gorm:"column:subject"`
	Issuer    string    `json:"issuer" gorm:"column:issuer"`
	TEndTime  time.Time `json:"t_end_time" gorm:"column:end_time"`
}

func GetSslConfig(webId int64) RSSLItem {
	var result RSSLItem
	sqlStr := `select * from ssl_config where web_id=@webId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"webId": webId,
	}).Scan(&result)
	return result
}

type RActiveSSLItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	WebId             string    `json:"web_id" gorm:"column:web_id"`
	Name              string    `json:"name" gorm:"column:name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           int64     `json:"group_id" gorm:"column:group_id"`
	Frequency         int       `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int       `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int       `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int       `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode int       `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	ApiHttpStatusCode string    `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64   `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64   `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64   `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int       `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int       `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	FailedWaitTimes   int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt         LocalTime `json:"created_at" gorm:"column:created_at"`
	Status            int       `json:"status" gorm:"column:status"`
}

func GetAllActiveSSLList() []RActiveSSLItem {
	var result []RActiveSSLItem
	sqlStr := `
	SELECT
	*
FROM
	ssl_config
LEFT JOIN web_list on web_list.id=ssl_config.web_id
LEFT JOIN monitor_policy ON monitor_policy.id = web_list.policy_id
WHERE
	ssl_config.deleted_at IS NULL`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
