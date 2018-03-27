package admin

import (
	"net/http"

	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/servers"
	"github.com/gorilla/mux"
)

// NewAPIServer .
func NewAPIServer(config *config.AdminAPIServerConfig, uiData string) servers.Server {
	r := mux.NewRouter()
	apiSubroute := r.PathPrefix("/api").Subrouter()

	messageHandler := &MessageHandler{}
	messageSubroute := apiSubroute.PathPrefix("/message").Subrouter()
	messageSubroute.Path("/{id}").HandlerFunc(messageHandler.View)
	messageSubroute.Path("").HandlerFunc(messageHandler.Messages)

	attachmentHandler := &AttachmentHandler{}
	attachmentSubroute := apiSubroute.PathPrefix("/attachment").Subrouter()
	attachmentSubroute.Path("/{id}/download").HandlerFunc(attachmentHandler.Download)

	userHandler := &UserHandler{}
	userSubroute := apiSubroute.PathPrefix("/user").Subrouter()
	userSubroute.Path("").HandlerFunc(userHandler.Search)

	if uiData == "" {
		uiData = "./frontend"
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(uiData)))

	return servers.NewHTTPServer(config.ListenAddress, r)
}
