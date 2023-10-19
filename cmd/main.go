package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"file-storage/config"
	"file-storage/internal/api"
	"file-storage/internal/repositories"
	"file-storage/internal/storage"
	"file-storage/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		utils.ErrorLogger.Fatalf("Failed to load configuration: %v", err)
	}

	// Инициализация серверов хранения
	var servers []repositories.StorageServer
	for _, dir := range cfg.StorageDirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			utils.ErrorLogger.Fatalf("Failed to init repos: %v", err)
		}
		servers = append(servers, storage.NewFileStorage(dir))
	}

	serverRepo := repositories.NewServerRepository(servers)
	fileHandler := delivery.NewFileHandler(serverRepo)

	http.HandleFunc("/upload", fileHandler.UploadFile)
	http.HandleFunc("/download", fileHandler.DownloadFile)
	http.HandleFunc("/status", delivery.StatusHandler)
	http.Handle("/metrics", promhttp.Handler())

	err = http.ListenAndServe(cfg.ServerPort, nil)
	if err != nil {
		utils.ErrorLogger.Fatalf("Failed to start server: %v", err)
	}
}
