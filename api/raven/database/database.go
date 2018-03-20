package database

import (
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/jinzhu/gorm"
	"github.com/palantir/stacktrace"
)

// Connection .
var Connection *gorm.DB

// Connect .
func Connect(config *config.DatabaseConfig) error {
	db, err := gorm.Open(config.Driver, config.ConnectionString)
	if err != nil {
		return stacktrace.Propagate(err, "cannot connect to database")
	}
	Connection = db
	return nil
}

// Close .
func Close() {
	if Connection != nil {
		Connection.Close()
	}
}
