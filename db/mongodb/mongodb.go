package mongodb

import (
	"os"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"gopkg.in/mgo.v2"
)

var defaultTimeout = time.Second * 5

// Settings holds all pertient connection parameters for MongoDB
var Settings *mongoDBSettings

type mongoDBSettings struct {
	Database string
	Host     string
	URL 	string
}

func init() {
	LoadConfig()
}

// LoadConfig (re)loads the influxDB config so a connection can be initialized
func LoadConfig() {
	database := os.Getenv("MONGODB")
	if database == "" {
		database = "bbqberry"
	}

	host := os.Getenv("MONGODB_HOST")
	if host == "" {
		host = "mongodb"
	}

	Settings = &mongoDBSettings{
		Host: host,
		Database: database,
	}
	log.Infof("action=LoadConfig mongoDBSettings=%+v", Settings )
}

// GetSession establishes a connection to the mongo database and returns the default database along with the session.
// The session must be closed by the caller.
func GetSession() (*mgo.Session, *mgo.Database, error){
	session, err := mgo.DialWithTimeout(Settings.Host, defaultTimeout)
	if ( err != nil ) {
		return nil, nil, err
	}
	db := session.DB("test")
	return session, db, nil
}