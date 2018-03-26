package mailgun

import (
	"net/http"

	"github.com/anduintransaction/raven/api/raven/utils"
	"github.com/sirupsen/logrus"
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
