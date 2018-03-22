package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/anduintransaction/raven/api/raven/database"
	"github.com/anduintransaction/raven/api/raven/model"
	"github.com/anduintransaction/raven/api/raven/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/palantir/stacktrace"
)

// MessageHandler .
type MessageHandler struct {
}

// View .
func (h *MessageHandler) View(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logrus.Warn("Invalid id: %q", mux.Vars(r)["id"])
		utils.ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	var email model.Email
	err = database.Connection.Where("id = ?", id).Preload("EmailContent").Preload("Attachments").First(&email).Error
	if err != nil {
		logrus.Error(stacktrace.Propagate(err, "cannot query email"))
		utils.ResponseServerError(w)
		return
	}
	err = utils.ResponseJSON(w, http.StatusOK, email)
	if err != nil {
		logrus.Error(err)
	}
}

// Messages .
func (h *MessageHandler) Messages(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var query []byte
	var err error
	if contentType == "application/json" {
		query, err = ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error(stacktrace.Propagate(err, "cannot parse form"))
			utils.ResponseServerError(w)
			return
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			logrus.Error(stacktrace.Propagate(err, "cannot parse form"))
			utils.ResponseServerError(w)
			return
		}
		query = []byte(r.Form.Get("query"))
	}
	messageQuery := &MessageQuery{
		Page:         1,
		ItemsPerPage: 10,
	}
	if len(query) > 0 {
		err = json.Unmarshal(query, messageQuery)
		if err != nil {
			logrus.Warn(stacktrace.Propagate(err, "cannot parse query: %s", string(query)))
			utils.ResponseError(w, http.StatusBadRequest, "invalid query")
			return
		}
	}
	db := h.buildQueryFilter(database.Connection, messageQuery)
	var count int
	err = db.Model(&model.Email{}).Count(&count).Error
	if err != nil {
		logrus.Error(stacktrace.Propagate(err, "cannot count"))
		utils.ResponseServerError(w)
		return
	}
	db = h.buildQueryPage(db, messageQuery)
	db = db.Preload("Attachments")
	var emails []*model.Email
	err = db.Find(&emails).Error
	if err != nil {
		logrus.Error(stacktrace.Propagate(err, "cannot query for email"))
		utils.ResponseServerError(w)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, &MessagesResponse{
		Count:  count,
		Emails: emails,
	})
}

func (h *MessageHandler) buildQueryFilter(db *gorm.DB, messageQuery *MessageQuery) *gorm.DB {
	if messageQuery.Filter != nil {
		fromFilter := strings.TrimSpace(messageQuery.Filter.From)
		toFilter := strings.TrimSpace(messageQuery.Filter.To)
		if fromFilter != "" {
			db = db.Where(&model.Email{FromEmail: fromFilter})
		}
		if toFilter != "" {
			db = db.Where(&model.Email{ToEmail: toFilter})
		}
	}
	search := strings.TrimSpace(messageQuery.Search)
	if search != "" {
		search = "%" + search + "%"
		db = db.
			Where("from_email ILIKE ?", search).
			Or("from_name ILIKE ?", search).
			Or("to_email ILIKE ?", search).
			Or("to_name ILIKE ?", search).
			Or("subject ILIKE ?", search)
	}
	return db
}

func (h *MessageHandler) buildQueryPage(db *gorm.DB, messageQuery *MessageQuery) *gorm.DB {
	for _, sorter := range messageQuery.Sorts {
		field := strings.TrimSpace(sorter.Field)
		direction := strings.ToLower(strings.TrimSpace(sorter.Direction))
		if direction != "desc" {
			direction = "asc"
		}
		if field != "" {
			db = db.Order(fmt.Sprintf("%s %s", field, direction))
		}
	}
	limit := messageQuery.ItemsPerPage
	offset := (messageQuery.Page - 1) * messageQuery.ItemsPerPage
	db = db.Offset(offset).Limit(limit)
	return db
}

// MessageQuery .
type MessageQuery struct {
	Filter       *MessageFilter
	Search       string
	Sorts        []*MessageSorter
	Page         int // Begin from 1, not 0
	ItemsPerPage int
}

// MessageFilter .
type MessageFilter struct {
	From string
	To   string
}

// MessageSorter .
type MessageSorter struct {
	Field     string
	Direction string
}

// MessagesResponse .
type MessagesResponse struct {
	Count  int
	Emails []*model.Email
}
