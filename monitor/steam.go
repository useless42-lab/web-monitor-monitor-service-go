package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ServerItem struct {
	Addr          string `json:"addr"`
	Gameport      int    `json:"gameport"`
	SteamId       string `json:"steamid"`
	Name          string `json:"name"`
	AppId         int    `json:"appid"`
	Gamedir       string `json:"gamedir"`
	Version       string `json:"version"`
	Product       string `json:"product"`
	Region        int    `json:"region"`
	PlayersOnline int    `json:"players"`
	PlayersMax    int    `json:"max_players"`
	Bots          int    `json:"bots"`
	Map           string `json:"map"`
	Secure        bool   `json:"secure"`
	Dedicated     bool   `json:"dedicated"`
	OS            string `json:"os"`
	GameType      string `json:"gametype"`
}
type ServersBlock struct {
	Servers []ServerItem `json:"servers"`
}
type ResponseItem struct {
	Response ServersBlock `json:"response"`
}

func MonitorSteam(key string, path string) (checkSuccess bool, name string, playersMax int, playersOnline int, elapsed int64) {
	t := time.Now()
	targetPath := `https://api.steampowered.com/IGameServersService/GetServerList/v1/?key=` + key + `&filter=addr\` + path
	req, _ := http.NewRequest("GET", targetPath, nil)
	resp, _ := http.DefaultClient.Do(req)

	var aa ResponseItem
	body, _ := ioutil.ReadAll(resp.Body)
	err1 := json.Unmarshal(body, &aa)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var checkSuccessBool bool = false

	elapsedTime := time.Since(t) / 1e6
	if len(aa.Response.Servers) > 0 {
		checkSuccessBool = true
		return checkSuccessBool, aa.Response.Servers[0].Name, aa.Response.Servers[0].PlayersMax, aa.Response.Servers[0].PlayersOnline, int64(elapsedTime)
	}
	return checkSuccessBool, "", 0, 0, 0
}
