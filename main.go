package main

import (
	"time"

	model "./model"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("cli:main")

const databaseRetryDuration = (time.Second * 2)
const databaseRetryAttempts = (5)

func main() {
	var spec Specification
	var db *gorm.DB
	var err error

	InitializeLogging()

	if err = envconfig.Process("dynchan", &spec); err != nil {
		log.Fatal("environment configuration error", err)
		return
	}

	attempt := 1

try_database_again:
	if db, err = gorm.Open(spec.DatabaseDialect, spec.DatabaseConnection); err != nil {
		log.Error("database connection error", err)

		if attempt < databaseRetryAttempts {
			attempt++
			log.Errorf("trying again in %v for attempt %v of %v", databaseRetryDuration, attempt, databaseRetryAttempts)

			time.Sleep(databaseRetryDuration)
			goto try_database_again
		} else {
			log.Fatalf("database connection failed after %v attempts", databaseRetryAttempts)
			return
		}
	}

	model.AutoMigrate(db)

}
