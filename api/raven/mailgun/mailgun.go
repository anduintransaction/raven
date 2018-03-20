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

	v3 := r.PathPrefix("/v3").Subrouter()
	messages := &MessageHandler{}
	v3.Methods("POST").Path("/{domain}/messages").HandlerFunc(messages.Send)
	return servers.NewHTTPServer(config.ListenAddress, r)
}
