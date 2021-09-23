package main

import (
	"fmt"
	"log"

	"github.com/rumblefrog/go-a2s"
)

type ApplicationInformation struct {
	CurrentPlayers    int       `json:"Current-Players"`
	MaxPlayers        int       `json:"Max-Players"`
	Map               string    `json:"Map,omitempty"`
	PasswordProtected bool      `json:"Password-Protected"`
	VAC               bool      `json:"VAC,omitempty"`
	Players           []*Player `json:"Player-Info,omitempty"`
}

func getApplicationInformation() *ApplicationInformation {
	switch config.QueryType {
	case "a2s":
		return a2sInfo()
	default:
		log.Printf("No query method specified.")
		return &ApplicationInformation{}
	}
}

func a2sQuery() (info *a2s.ServerInfo, players *a2s.PlayerInfo, err error) {
	log.Printf("Creating a2s client")
	address := fmt.Sprintf("application:%v", config.QueryPort)
	client, err := a2s.NewClient(address)
	if err == nil {
		defer client.Close()
		log.Printf("Querying a2s info.")
		info, err = client.QueryInfo()
		players, err = client.QueryPlayer()
		if err != nil {
			log.Printf("Error querying info: %v", err)
			return info, players, err
		}
		client.Close()
		return info, players, nil
	} else {
		log.Printf("Error creating client. Error: %v", err)
		return info, players, err
	}
}

func a2sInfo() *ApplicationInformation {
	appInfo := &ApplicationInformation{}
	info, players, err := a2sQuery()
	if err != nil {
		log.Printf("Error getting info: %v", err)
	} else {
		if info != nil {
			appInfo.CurrentPlayers = int(info.Players)
			appInfo.Map = info.Map
			appInfo.PasswordProtected = info.Visibility
			appInfo.MaxPlayers = int(info.MaxPlayers)
			appInfo.VAC = info.VAC
		}
		if players != nil {
			var playerInfo []*Player
			for _, player := range players.Players {
				newPlayer := &Player{
					Name:     player.Name,
					Duration: player.Duration,
				}
				playerInfo = append(playerInfo, newPlayer)
			}
			appInfo.Players = playerInfo
		}
	}
	return appInfo
}
