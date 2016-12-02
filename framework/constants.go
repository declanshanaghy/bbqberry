package framework

func init() {
	Constants = constants{
		ServiceName: "bbqberry",
		Version:     "v1",
	}
}

type constants struct {
	ServiceName string
	Version     string
}

// Constants contains static information about the running service
var Constants constants
