package database

import (
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
	db, err := gorm.Open(config.Driver, config.ConnectionString)
	if err != nil {
		return stacktrace.Propagate(err, "cannot connect to database")
	}
	defer db.Close()
	return stacktrace.Propagate(db.AutoMigrate(&model.Email{}, &model.Attachment{}).Error, "cannot migrate database")
}
