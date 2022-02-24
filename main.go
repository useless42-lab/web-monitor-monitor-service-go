package main

import (
	"WebMonitor/jobs"
	"fmt"

	"github.com/bamzi/jobrunner"
)

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 15s", jobs.CheckWebJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckServerJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckApiJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckTcpJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckDnsJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckSteamGameServerJob{})
	jobrunner.Schedule("@every 15s", jobs.CheckMinecraftServerJob{})
	jobrunner.Schedule("@daily", jobs.CheckPlanJob{})
	jobrunner.Schedule("@weekly", jobs.DeleteLogJob{})

	var str string
	fmt.Scan(&str)
}
