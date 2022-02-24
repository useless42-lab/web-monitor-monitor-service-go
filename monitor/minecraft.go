package monitor

import (
	"time"

	"github.com/PassTheMayo/mcstatus/v3"
)

func MonitorMinecraft(platformVersion int, path string) (checkSuccess bool, max int, online int, versionName string, elapsed time.Duration) {
	if platformVersion == 1 {
		// java
		respJava, err := mcstatus.Status(path, 25565, mcstatus.JavaStatusOptions{Timeout: time.Duration(3) * time.Second})
		if err != nil {
			return false, 0, 0, "", 0
		} else {
			// fmt.Println(respJava.MOTD)
			return true, respJava.Players.Max, respJava.Players.Online, respJava.Version.Name, respJava.Latency
		}
	} else {
		// 基岩
		respBedrock, err := mcstatus.StatusBedrock(path, 19132, mcstatus.BedrockStatusOptions{Timeout: time.Duration(3) * time.Second})
		if err != nil {
			return false, 0, 0, "", 0
		} else {
			return true, int(*respBedrock.MaxPlayers), int(*respBedrock.OnlinePlayers), *respBedrock.Version, 0
		}
	}
}
