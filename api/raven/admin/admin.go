package admin

import (
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/servers"
	"github.com/gorilla/mux"
)

// NewAPIServer .
func NewAPIServer(config *config.AdminAPIServerConfig) servers.Server {
	r := mux.NewRouter()
	return servers.NewHTTPServer(config.ListenAddress, r)
}
