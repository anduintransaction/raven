package admin

import (
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/servers"
	"github.com/gorilla/mux"
)

// NewAPIServer .
func NewAPIServer(config *config.AdminAPIServerConfig) servers.Server {
	r := mux.NewRouter()
	apiSubroute := r.PathPrefix("/api").Subrouter()

	messageHandler := &MessageHandler{}
	messageSubroute := apiSubroute.PathPrefix("/message").Subrouter()
	messageSubroute.Path("/{id}").HandlerFunc(messageHandler.View)
	messageSubroute.Path("").HandlerFunc(messageHandler.Messages)

	return servers.NewHTTPServer(config.ListenAddress, r)
}
