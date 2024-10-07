package config

import "os"

type Config struct {
	AppName string
	Port    string
	DBUrl   string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		AppName: "VisionLens Multimodal Computer Vision Analyzer",
		Port:    port,
		DBUrl:   os.Getenv("DATABASE_URL"),
	}
}
