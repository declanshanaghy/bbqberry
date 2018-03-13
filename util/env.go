package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Polarishq/middleware/framework/log"
)

// GetEnvInt retrieves the given environment variable as an int and
// returns the given default if it is not set or a parsing error occurs
func GetEnvInt(key string, def int64) (r int64, err error) {
	s := os.Getenv("DB_TIMEOUT_MILLIS")
	r = def
	if s != "" {
		r, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			err = fmt.Errorf("Error parsing %s from %s, reverting to default %d", s, key, def)
			r = def
		}
	}
	return
}

// GetEnvMillisAsDuration attempts to retrieve the given environment variable represented in milliseconds
// and parse it into a time.Duration. If the variable is not set the given default is returned.
// If an error occurs it is logged and the default is returned
func GetEnvMillisAsDuration(key string, def int64) time.Duration {
	timeoutMillis, err := GetEnvInt(key, def)
	if err != nil {
		log.Error(err)
	}
	return time.Millisecond * time.Duration(timeoutMillis)
}
