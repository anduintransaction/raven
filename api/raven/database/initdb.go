package database

import (
	"time"

	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/model"
	"github.com/jinzhu/gorm"
	"github.com/palantir/stacktrace"

	// Blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB .
func InitDB(config *config.DatabaseConfig) error {
	var db *gorm.DB
	var err error
	maxRetry := 10
	for i := 0; i < maxRetry; i++ {
		db, err = gorm.Open(config.Driver, config.ConnectionString)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		return stacktrace.Propagate(err, "cannot connect to database")
	}
	defer db.Close()
	return stacktrace.Propagate(db.AutoMigrate(&model.Message{}, &model.Email{}, &model.EmailContent{}, &model.Attachment{}, &model.AttachmentData{}).Error, "cannot migrate database")
}
