package main

import (
	"github.com/2flow/gokies/httputils"
	filestorage "github.com/2flow/gokies/storageabstraction/azureblobs"
	"github.com/go-kit/log"
	"os"
)

func main() {
	azureAccountName := os.Getenv("AZURE_ACCOUNT_NAME")
	azureAccountKey := os.Getenv("AZURE_ACCOUNT_KEY")
	azureContainerName := os.Getenv("AZURE_CONTAINER_NAME")
	azureStorageURL := os.Getenv("AZURE_STORAGE_URL")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	httpServer := httputils.NewHTTPServer(logger)
	storage := filestorage.NewAzureStorage(azureAccountName,
		azureAccountKey,
		azureContainerName,
		azureStorageURL)

	fileContainer := httputils.HTTPFileContainer{
		FileStorage: storage,
		RootDir:     "/",
	}

	httpServer.SetRoutes(fileContainer.ProvideFileHandler())
	httpServer.StartHTTPSServer()
	httpServer.StartHTTPServer()
}
