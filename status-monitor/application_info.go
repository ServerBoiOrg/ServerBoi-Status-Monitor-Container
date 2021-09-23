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
		// Put all this in A2S Query method
		a2sInfo, a2sPlayers, err := a2sQuery()
		if err != nil {
			log.Printf("Error getting a2s info: %v", err)
		}
		var playerInfo []*Player
		if a2sPlayers != nil {
			for _, player := range a2sPlayers.Players {
				newPlayer := &Player{
					Name:     player.Name,
					Duration: player.Duration,
				}
				playerInfo = append(playerInfo, newPlayer)
			}
		}
		return &ApplicationInformation{
			CurrentPlayers:    int(a2sInfo.Players),
			MaxPlayers:        int(a2sInfo.MaxPlayers),
			Map:               a2sInfo.Map,
			PasswordProtected: a2sInfo.Visibility,
			VAC:               a2sInfo.VAC,
			Players:           playerInfo,
		}
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
		if err != nil {
			log.Printf("Error querying info: %v", err)
		}
		players, err = client.QueryPlayer()
		if err != nil {
			log.Printf("Error querying players: %v", err)
		}
		client.Close()
		return info, players, nil
	} else {
		log.Printf("Error creating client. Error: %v", err)
		return info, players, err
	}
}
