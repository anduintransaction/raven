package admin

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/anduintransaction/raven/api/raven/database"
	"github.com/anduintransaction/raven/api/raven/utils"
	"github.com/jinzhu/gorm"
	"github.com/palantir/stacktrace"
)

// UserHandler .
type UserHandler struct {
}

// Search .
func (h *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	query, err := getQuery(r)
	if err != nil {
		logrus.Error(err)
		utils.ResponseServerError(w)
		return
	}
	userQuery := &UserQuery{
		Query: "",
	}
	if len(query) > 0 {
		err = json.Unmarshal(query, userQuery)
		if err != nil {
			logrus.Error(stacktrace.Propagate(err, "cannot parse query: %s", string(query)))
			utils.ResponseError(w, http.StatusBadRequest, "invalid query")
			return
		}
	}
	searchPhase := strings.TrimSpace(userQuery.Query)
	fromQuery := database.
		Connection.
		Table("emails").
		Select("DISTINCT(from_email)")
	if searchPhase != "" {
		fromQuery = fromQuery.
			Where("from_email ILIKE ?", "%"+searchPhase+"%").
			Or("from_name ILIKE ?", "%"+searchPhase+"%")
	}
	fromUsers, err := h.queryUsers(fromQuery)
	if err != nil {
		logrus.Error(err, "cannot query from users")
		utils.ResponseServerError(w)
		return
	}
	toQuery := database.
		Connection.
		Table("emails").
		Select("DISTINCT(to_email)")
	if searchPhase != "" {
		toQuery = toQuery.
			Where("to_email ILIKE ?", "%"+searchPhase+"%").
			Or("to_name ILIKE ?", "%"+searchPhase+"%")
	}
	toUsers, err := h.queryUsers(toQuery)
	if err != nil {
		logrus.Error(err, "cannot query from users")
		utils.ResponseServerError(w)
		return
	}
	allUsers := h.mergeUsers(fromUsers, toUsers)
	err = utils.ResponseJSON(w, http.StatusOK, &UserResponse{
		Count: len(allUsers),
		Users: allUsers,
	})
	if err != nil {
		logrus.Error(err)
	}
}

func (h *UserHandler) makeUserQuery(db *gorm.DB, userQuery *UserQuery) *gorm.DB {
	searchPhase := strings.TrimSpace(userQuery.Query)
	if searchPhase != "" {
		searchPhase = "%" + searchPhase + "%"
		return db.
			Where("from_email ILIKE ?", searchPhase).
			Or("from_name ILIKE ?", searchPhase)
	}
	return db
}

func (h *UserHandler) queryUsers(db *gorm.DB) ([]string, error) {
	users := []string{}
	rows, err := db.Rows()
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot query users")
	}
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			return nil, stacktrace.Propagate(err, "cannot scan users")
		}
		users = append(users, user)
	}
	return users, nil
}

func (h *UserHandler) mergeUsers(l1, l2 []string) []string {
	s := make(map[string]struct{})
	for _, u := range l1 {
		s[u] = struct{}{}
	}
	for _, u := range l2 {
		s[u] = struct{}{}
	}
	m := []string{}
	for u := range s {
		m = append(m, u)
	}
	sort.Strings(m)
	return m
}

// UserQuery .
type UserQuery struct {
	Query string
}

// UserResponse .
type UserResponse struct {
	Count int
	Users []string
}
