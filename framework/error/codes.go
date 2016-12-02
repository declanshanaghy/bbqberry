package error_codes

const (
	ErrInfluxUnavailable = int32(900)
	ErrInfluxWrite = int32(910)
)

var messages = map[int32]string{
	ErrInfluxUnavailable: "Unable to initialize connection to InfluxDB server",
	ErrInfluxWrite: "An error occurred writing data to influxdb",
}

/*
Returns a text message for the given error code.
It returns the empty string if the code is unknown.
 */
func GetText(code int32) string {
	return messages[code]
}