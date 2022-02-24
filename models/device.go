package models

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func CheckMonitorRegion(regionStr string) bool {
	regionArr := strings.Split(regionStr, ",")
	var isExist bool = false
	if len(regionArr) > 0 {
		for _, item := range regionArr {
			if item == os.Getenv("REGION") {
				isExist = true
			}
		}
	} else {
	}
	return isExist
}

type RSimpleDeviceIdItem struct {
	Id         string `json:"id"`
	DeviceType int    `json:"device_type"`
}

/*
获取该用户创建团队下所有设备编号
*/
func GetAllDeviceListOwnerByUserId(userId int64) []RSimpleDeviceIdItem {
	sqlStr := `
	SELECT
			id,1 as device_type
		FROM
			web_list AS wl
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
				AND deleted_at IS NULL
				AND status=1
			)
		UNION ALL
			SELECT
				id,2 as device_type
			FROM
				server_list AS sl
			WHERE
				group_id IN (
					SELECT
						id
					FROM
						device_group
					WHERE
						team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
					AND deleted_at IS NULL
					AND status=1
				)
			UNION ALL
				SELECT
					id,3 as device_type
				FROM
					api_list AS al
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
UNION ALL
				SELECT
					id,4 as device_type
				FROM
					tcp_list AS tl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
UNION ALL
				SELECT
					id,5 as device_type
				FROM
					dns_list AS dl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
UNION ALL
				SELECT
					id,6 as device_type
				FROM
					heartbeat_list AS hbl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
UNION ALL
				SELECT
					id,7 as device_type
				FROM
					steam_game_server_list AS sgsl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
UNION ALL
				SELECT
					id,8 as device_type
				FROM
					minecraft_server_list AS msl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null and role=2)
						AND deleted_at IS NULL
						AND status=1
					)
	`
	var result []RSimpleDeviceIdItem
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}
func FilterDevice(deviceType int) string {
	var device string
	if deviceType == 1 {
		device = "web_list"
	}
	if deviceType == 2 {
		device = "server_list"
	}
	if deviceType == 3 {
		device = "api_list"
	}
	if deviceType == 4 {
		device = "tcp_list"
	}
	if deviceType == 5 {
		device = "dns_list"
	}
	if deviceType == 6 {
		device = "heartbeat_list"
	}
	if deviceType == 7 {
		device = "steam_game_server_list"
	}
	if deviceType == 8 {
		device = "minecraft_server_list"
	}
	return device
}
func ResetDeviceOwnerByUser(deviceId int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `
	update ` + device + ` set status=3 where id=@deviceId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
	})
}
