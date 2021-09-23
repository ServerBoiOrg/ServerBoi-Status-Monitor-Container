package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	IP           string
	Created      string
	ClientPort   int
	QueryPort    int
	QueryType    string
	Application  string
	Name         string
	ID           string
	OwnerID      string
	OwnerName    string
	HostOS       string
	Architecture string
	Provider     string
	HardwareType string
	Region       string
}

func getIP() string {
	resp, err := http.Get("http://checkip.amazonaws.com")
	if err == nil {
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		return strings.TrimSpace(string(b))
	} else {
		return ""
	}
}

func getConfig() *Config {
	env := godotenv.Load(".env")
	if env == nil {
		log.Fatalf("Error loading .env file")
	}

	clientPort, _ := strconv.Atoi(os.Getenv("CLIENT_PORT"))
	queryPort, _ := strconv.Atoi(os.Getenv("QUERY_PORT"))

	return &Config{
		IP:           getIP(),
		Created:      os.Getenv("CREATED"),
		ClientPort:   clientPort,
		QueryPort:    queryPort,
		QueryType:    os.Getenv("QUERY_TYPE"),
		Application:  os.Getenv("APPLICATION"),
		Name:         os.Getenv("NAME"),
		ID:           os.Getenv("ID"),
		OwnerID:      os.Getenv("OWNER_ID"),
		OwnerName:    os.Getenv("OWNER_NAME"),
		HostOS:       os.Getenv("HOST_OS"),
		Architecture: os.Getenv("ARCHITECTURE"),
		Provider:     os.Getenv("PROVIDER"),
		HardwareType: os.Getenv("HARDWARE_TYPE"),
		Region:       os.Getenv("REGION"),
	}
}
