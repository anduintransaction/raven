package mailgun

import (
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/servers"
	"github.com/gorilla/mux"
)

// NewAPIServer .
func NewAPIServer(config *config.MailgunAPIServerConfig) servers.Server {
	r := mux.NewRouter()
	home := &HomeHandler{}
	r.Path("/").HandlerFunc(home.Home)
	return servers.NewHTTPServer(config.ListenAddress, r)
}
