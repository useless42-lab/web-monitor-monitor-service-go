package models

func DeleteWebLog() {
	sqlStr := `delete From web_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}

func DeleteServerLog() {
	sqlStr := `delete From server_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteApiLog() {
	sqlStr := `delete From api_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteTcpLog() {
	sqlStr := `delete From tcp_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteDnsLog() {
	sqlStr := `delete From dns_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteHeartbeatLog() {
	sqlStr := `delete From heartbeat_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteSteamGameServerLog() {
	sqlStr := `delete From steam_game_server_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
func DeleteMinecraftServerLog() {
	sqlStr := `delete From minecraft_server_log where DATE(created_at) <= DATE(DATE_SUB(NOW(),INTERVAL 180 day))`
	DB.Exec(sqlStr)
}
