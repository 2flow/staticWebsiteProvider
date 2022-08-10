package httputils

import (
	"github.com/2flow/gokies/storageabstraction"
	"io"
	"net/http"
	"strings"
)

type HTTPFileContainer struct {
	FileStorage storageabstraction.IFileStorage
	RootDir     string
}

func (container HTTPFileContainer) ProvideFileHandler() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		fetchDest := request.Header.Get("Sec-Fetch-Dest") // if index.html --> document otherwise script
		routePath := request.URL.Path

		// if the document is requested return the index.html
		// this should work
		if fetchDest == "document" {
			routePath = "/index.html"
		} else if fetchDest == "" {
			parts := strings.Split(routePath, ".")
			if len(parts) == 1 {
				routePath = "/index.html"
			}
		} else if routePath == "/" {
			routePath = "/index.html"
		}

		if strings.HasPrefix(routePath, "/") && container.RootDir != "/" {
			routePath = container.RootDir + routePath
		}
		if strings.HasPrefix(routePath, "/") {
			routePath = routePath[1:]
		}

		reader, err := container.FileStorage.DownloadFile(routePath)

		if err != nil {
			HTTPRoutingErrorHandler("Unable to read file", err).EncodeStatus(responseWriter, http.StatusInternalServerError)
			return
		}

		defer reader.Close()
		SetContentType(responseWriter, reader, routePath)
		responseWriter.WriteHeader(http.StatusOK)
		io.Copy(responseWriter, reader)
	})
}
