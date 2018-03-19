package mailgun

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/anduintransaction/raven/api/raven/utils"
)

// HomeHandler .
type HomeHandler struct {
}

// Home .
func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	err := utils.ResponseJSON(w, http.StatusOK, struct {
		Version string `json:"version"`
	}{
		Version: utils.Version,
	})
	if err != nil {
		logrus.Error(err)
	}
}
