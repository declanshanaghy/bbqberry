package mongodb

import (
	"os"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/util"
	"gopkg.in/mgo.v2"
)

// Settings holds all pertient connection parameters for MongoDB
var Settings *mongoDBSettings

type mongoDBSettings struct {
	Database string
	Host     string
	Timeout  time.Duration
}

func init() {
	LoadConfig()
}

// LoadConfig (re)loads the influxDB config so a connection can be initialized
func LoadConfig() {
	database := os.Getenv("MONGODB")
	if database == "" {
		database = framework.DefaultDB
	}

	host := os.Getenv("MONGODB_HOST")
	if host == "" {
		host = "mongodb"
	}

	timeout := util.GetEnvMillisAsDuration("DB_TIMEOUT_MILLIS", 5000)
	Settings = &mongoDBSettings{
		Host:     host,
		Database: database,
		Timeout:  timeout,
	}
	log.Infof("action=LoadConfig mongoDBSettings=%+v", Settings)
}

// newSession establishes a connection to the mongo database and returns the default database along with the session.
// The session must be closed by the caller.
func newSession() (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.DialWithTimeout(Settings.Host, Settings.Timeout)
	if err != nil {
		return nil, nil, err
	}
	db := session.DB(Settings.Database)
	return session, db, nil
}
