package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controllers "github.com/jlettman/dynchan/controllers"
	globals "github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
	"github.com/kelseyhightower/envconfig"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("cli:main")

const databaseRetryDuration = (time.Second * 2)
const databaseRetryAttempts = (5)

func main() {
	var spec Specification
	var err error

	InitializeLogging()

	if err = envconfig.Process("dynchan", &spec); err != nil {
		log.Fatal("environment configuration error", err)
		return
	}

	attempt := 1

try_database_again:
	if globals.DB, err = gorm.Open(spec.DatabaseDialect, spec.DatabaseConnection); err != nil {
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

	globals.DB.SetLogger(GormLogging{})
	models.AutoMigrate(globals.DB)

	router := gin.New()
	router.Use(gin.Recovery())

	controllers.Routes(router)
	router.Run(":8080")

}
