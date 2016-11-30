package framework

func init() {
	ConstantsObj = Constants{
		ServiceName:    "bbqberry",
		Version:		"v1",
	}
}

type Constants struct {
	ServiceName    string
	Version        string
}

var ConstantsObj Constants