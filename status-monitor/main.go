package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/rumblefrog/go-a2s"
)

var (
	statusFile = "server_status.json"
	config     = getConfig()
	address    = fmt.Sprintf("%v:%v", config.IP, config.Port)
)

func main() {
	log.Printf("Starting")
	info, players := waitForClientStart()
	writeServerInfo(info, players)

	router := configureRouter()

	go continouslyUpdateStatus()
	router.Run(":7032")
}

func waitForClientStart() (a2sInfo *a2s.ServerInfo, a2sPlayers *a2s.PlayerInfo) {
	log.Printf("Waiting for client to start")

	for {
		log.Printf("Querying address %v", address)
		info, players, err := a2sQuery(address)
		if err == nil {
			if info != nil {
				log.Printf("Client started")
				a2sInfo = info
				a2sPlayers = players
				break
			} else {

			}
		} else {
			log.Printf("No response, trying again in 30 seconds")
			time.Sleep(time.Duration(30) * time.Second)
		}
	}
	return a2sInfo, a2sPlayers
}

func continouslyUpdateStatus() {
	for {
		log.Printf("Updating server status")
		address := fmt.Sprintf("%v:%v", config.IP, config.Port)
		info, players, _ := a2sQuery(address)
		writeServerInfo(info, players)
		time.Sleep(time.Duration(30) * time.Second)
	}
}

type Player struct {
	Name     string  `json:"Name"`
	Duration float32 `json:"Duration"`
}

type ServerInfo struct {
	LastUpdate        string    `json:"LastUpdate"`
	Name              string    `json:"Name"`
	Application       string    `json:"Application"`
	ServerType        string    `json:"ServerType"`
	OS                string    `json:"OS"`
	PlayerCount       int       `json:"PlayerCount"`
	MaxPlayers        int       `json:"MaxPlayers"`
	Map               string    `json:"Map,omitempty"`
	PasswordProtected bool      `json:"PasswordProtected"`
	VAC               bool      `json:"VAC"`
	IP                string    `json:"IP"`
	Port              int       `json:"Port"`
	Players           []*Player `json:"Players"`
}

func writeServerInfo(info *a2s.ServerInfo, players *a2s.PlayerInfo) {
	log.Printf("Writing server information")
	port, _ := strconv.Atoi(config.Port)
	var playerInfo []*Player
	if players != nil {
		for _, player := range players.Players {
			newPlayer := &Player{
				Name:     player.Name,
				Duration: player.Duration,
			}
			playerInfo = append(playerInfo, newPlayer)
		}
	}

	serverInfo := ServerInfo{
		LastUpdate:        time.Now().UTC().String(),
		Name:              info.Name,
		Application:       info.Game,
		ServerType:        info.ServerType.String(),
		OS:                info.ServerOS.String(),
		PlayerCount:       int(info.Players),
		MaxPlayers:        int(info.MaxPlayers),
		Map:               info.Map,
		PasswordProtected: info.Visibility,
		VAC:               info.VAC,
		IP:                config.IP,
		Port:              port,
		Players:           playerInfo,
	}
	file, _ := json.MarshalIndent(serverInfo, "", " ")
	_ = ioutil.WriteFile(statusFile, file, 0644)
}

func a2sQuery(address string) (info *a2s.ServerInfo, players *a2s.PlayerInfo, err error) {
	client, err := a2s.NewClient(address)
	if err == nil {
		defer client.Close()
		info, _ = client.QueryInfo()
		players, _ = client.QueryPlayer()
		client.Close()
		return info, players, nil
	}
	return info, players, err
}
