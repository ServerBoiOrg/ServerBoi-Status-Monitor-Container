package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

var (
	statusFile = "server_status.json"
	config     *Config
)

func init() {
	config = getConfig()
}

func main() {
	log.Printf("Starting")
	router := configureRouter()
	go continouslyUpdateStatus()
	router.Run(":7032")
}

func continouslyUpdateStatus() {
	for {
		updateServerInfo()
		time.Sleep(time.Duration(30) * time.Second)
	}
}

func updateServerInfo() {
	log.Printf("Updating server information")
	updatedInfo := ServerInfo{
		General:           getGeneralInfo(),
		AppInfo:           getApplicationInformation(),
		ServiceInfo:       getServiceInfo(),
		SystemInformation: getSystemInfo(),
	}
	file, _ := json.MarshalIndent(updatedInfo, "", " ")
	_ = ioutil.WriteFile(statusFile, file, 0644)
}
