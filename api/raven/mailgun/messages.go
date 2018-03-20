package mailgun

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/mail"
	"net/url"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/anduintransaction/raven/api/raven/model"
	"github.com/anduintransaction/raven/api/raven/utils"
	"github.com/palantir/stacktrace"
)

// MessageHandler .
type MessageHandler struct {
}

// Send .
func (h *MessageHandler) Send(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var values url.Values
	var files map[string][]*multipart.FileHeader
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			logrus.Error(err)
			utils.ResponseServerError(w)
			return
		}
		values = r.Form
		files = nil
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		err := r.ParseMultipartForm(10 * 1024 * 1024)
		if err != nil {
			logrus.Error(err)
			utils.ResponseServerError(w)
			return
		}
		values = url.Values(r.MultipartForm.Value)
		files = r.MultipartForm.File
	} else {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid content type")
		return
	}
	fmt.Println(values.Get("to"))
	fmt.Println(files)

	response := &struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}{
		ID:      "ahihi",
		Message: "ahuhu",
	}
	err := utils.ResponseJSON(w, http.StatusOK, response)
	if err != nil {
		logrus.Error(err)
	}
}

func (h *MessageHandler) parseEmail(formValues url.Values, files map[string][]*multipart.FileHeader) ([]*model.Email, []*model.Attachment, error) {
	fromAddress, err := mail.ParseAddress(formValues.Get("from"))
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "cannot parse from address %q", formValues.Get("from"))
	}
	rcptAddresses, err := mail.ParseAddressList(formValues.Get("to"))
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "cannot parse rcpt address %q", formValues.Get("to"))
	}
	replyTo := formValues.Get("h:Reply-To")
	subject := formValues.Get("subject")
	htmlContent := formValues.Get("html")
	attachments := []*model.Attachment{}
	for _, attachmentHeader := range files["attachment"] {
		f, err := attachmentHeader.Open()
		if err != nil {
			return nil, nil, stacktrace.Propagate(err, "cannot open attachment %q", attachmentHeader.Filename)
		}
		defer f.Close()
		attachmentContent, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, nil, stacktrace.Propagate(err, "cannot read attachment %q", attachmentHeader.Filename)
		}
		attachment := &model.Attachment{
			Filename: attachmentHeader.Filename,
			Filemime: attachmentHeader.Header.Get("Content-Type"),
			Filesize: attachmentHeader.Size,
			Content:  attachmentContent,
		}
		attachments = append(attachments, attachment)
	}
	emails := []*model.Email{}
	for _, rcptAddress := range rcptAddresses {
		email := &model.Email{
			From:     fromAddress.Address,
			FromName: fromAddress.Name,
			To:       rcptAddress.Address,
			ToName:   rcptAddress.Name,
			RCPT:     formValues.Get("to"),
			ReplyTo:  replyTo,
			Subject:  subject,
			HTML:     htmlContent,
		}
		emails = append(emails, email)
	}
	return emails, attachments, nil
}
