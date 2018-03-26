package admin

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anduintransaction/raven/api/raven/database"
	"github.com/anduintransaction/raven/api/raven/model"
	"github.com/anduintransaction/raven/api/raven/utils"
	"github.com/gorilla/mux"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

// AttachmentHandler .
type AttachmentHandler struct {
}

// Download .
func (h *AttachmentHandler) Download(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logrus.Warn("Invalid id: %q", mux.Vars(r)["id"])
		utils.ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	var attachment model.Attachment
	err = database.Connection.Where("id = ?", id).Preload("AttachmentData").First(&attachment).Error
	if err != nil {
		logrus.Error(stacktrace.Propagate(err, "cannot query attachment"))
		utils.ResponseServerError(w)
		return
	}
	content := bytes.NewReader(attachment.AttachmentData.Content)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", attachment.Filename))
	http.ServeContent(w, r, attachment.Filename, attachment.CreatedAt, content)
}
