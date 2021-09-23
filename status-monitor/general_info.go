package main

import "time"

type Player struct {
	Name     string  `json:"Name"`
	Duration float32 `json:"Duration,omitempty"`
}

type ServerInfo struct {
	General           *GeneralInformation     `json:"General"`
	AppInfo           *ApplicationInformation `json:"Application"`
	ServiceInfo       *ServiceInformation     `json:"Service"`
	SystemInformation *SystemInformation      `json:"System"`
}

type GeneralInformation struct {
	Application string `json:"Application"`
	LastUpdated string `json:"Last-Updated"`
	Created     string `json:"Created"`
	Name        string `json:"Name"`
	ID          string `json:"ID"`
	OwnerID     string `json:"Owner-ID"`
	OwnerName   string `json:"Owner-Name"`
	ClientPort  int    `json:"Client-Port"`
	QueryPort   int    `json:"Query-Port,omitempty"`
	QueryType   string `json:"Query-Type,omityempty"`
	IP          string `json:"IP"`
	HostOS      string `json:"Host-OS"`
}

type ServiceInformation struct {
	Provider     string `json:"Provider"`
	HardwareType string `json:"Hardware-Type"`
	Region       string `json:"Region"`
}

func getGeneralInfo() *GeneralInformation {
	return &GeneralInformation{
		Application: config.Application,
		LastUpdated: time.Now().UTC().String(),
		Created:     config.Created,
		Name:        config.Name,
		ID:          config.ID,
		OwnerID:     config.OwnerID,
		OwnerName:   config.OwnerName,
		ClientPort:  config.ClientPort,
		QueryPort:   config.QueryPort,
		QueryType:   config.QueryType,
		IP:          config.IP,
		HostOS:      config.HostOS,
	}
}

func getServiceInfo() *ServiceInformation {
	return &ServiceInformation{
		Provider:     config.Provider,
		HardwareType: config.HardwareType,
		Region:       config.Region,
	}
}
