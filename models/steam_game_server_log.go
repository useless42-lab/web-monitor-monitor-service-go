package models

import "time"

func GetMonitorSteamGameServerSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM steam_game_server_log WHERE steam_game_server_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}

type SteamGameServerLogItem struct {
	DefaultModel
	SteamGameServerId string `json:"steam_game_server_id" gorm:"column:steam_game_server_id"`
	Name              string `json:"name" gorm:"column:name`
	PlayersMax        int    `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int    `json:"players_online" gorm:"column:players_online"`
	Elapsed           int    `json:"elapsed" gorm:"column:elapsed"`
	CheckSuccess      int    `json:"check_success" gorm:"column:check_success"`
	Region            string `json:"region" gorm:"column:region"`
}

type RSteamGameServerLogItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	SteamGameServerId string    `json:"steam_game_server_id" gorm:"column:steam_game_server_id"`
	Name              string    `json:"name" gorm:"column:name`
	PlayersMax        int       `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int       `json:"players_online" gorm:"column:players_online"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	CheckSuccess      int       `json:"check_success" gorm:"column:check_success"`
	Region            string    `json:"region" gorm:"column:region"`
}

func AddSteamGameServerLog(steamGameServerId string, name string, playersMax int, playersOnline int, elapsed int, checkSuccess int, region string) {
	data := SteamGameServerLogItem{
		SteamGameServerId: steamGameServerId,
		Name:              name,
		PlayersMax:        playersMax,
		PlayersOnline:     playersOnline,
		Elapsed:           elapsed,
		CheckSuccess:      checkSuccess,
		Region:            region,
	}
	DB.Table("steam_game_server_log").Create(&data)
}

func GetLatestSteamGameServerLog(steamGameServerId string, region string) RSteamGameServerLogItem {
	var result RSteamGameServerLogItem
	sqlStr := `select * from steam_game_server_log where steam_game_server_id=@steamGameServerId and region=@region order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"steamGameServerId": steamGameServerId,
		"region":            region,
	}).Scan(&result)
	return result
}
