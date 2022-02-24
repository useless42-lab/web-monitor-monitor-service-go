package models

import "time"

type DnsLogItem struct {
	DefaultModel
	DnsId        string `json:"dns_id" gorm:"column:dns_id"`
	DnsType      int    `json:"dns_type" gorm:"column:dns_type"`
	Elapsed      int64  `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int    `json:"check_success" gorm:"column:check_success"`
	Region       string `json:"region" gorm:"column:region"`
}

type RDnsLogItem struct {
	DnsId        string    `json:"dns_id" gorm:"column:dns_id"`
	DnsType      int       `json:"dns_type" gorm:"column:dns_type"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"column:check_success"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	Region       string    `json:"region" gorm:"column:region"`
}

func AddDnsLog(dnsId string, dnsType int, elapsed int64, responseData string, checkSuccess int, region string) {
	data := DnsLogItem{
		DnsId:        dnsId,
		DnsType:      dnsType,
		Elapsed:      elapsed,
		ResponseData: responseData,
		CheckSuccess: checkSuccess,
		Region:       region,
	}
	DB.Table("dns_log").Create(&data)
}

func GetLatestDnsLog(dnsId string, region string) RDnsLogItem {
	var result RDnsLogItem
	sqlStr := `select * from dns_log where dns_id=@dnsId and region=@region order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"dnsId":  dnsId,
		"region": region,
	}).Scan(&result)
	return result
}

func GetMonitorDnsSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM dns_log WHERE dns_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}
