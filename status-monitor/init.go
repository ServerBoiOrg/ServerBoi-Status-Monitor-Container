package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	IP          string
	Port        string
	QueryMethod string
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

	return &Config{
		getIP(),
		os.Getenv("PORT"),
		os.Getenv("QUERY_METHOD"),
	}
}
