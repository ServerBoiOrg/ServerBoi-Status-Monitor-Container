package main

import (
	sq "serverquery"
	"time"
)

func getGeneralInfo() *sq.GeneralInformation {
	return &sq.GeneralInformation{
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

func getServiceInfo() *sq.ServiceInformation {
	return &sq.ServiceInformation{
		Provider:     config.Provider,
		HardwareType: config.HardwareType,
		Region:       config.Region,
	}
}
