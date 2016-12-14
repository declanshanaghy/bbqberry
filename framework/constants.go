package framework

import (
	"os"
)

func init() {
	stub := false
	if os.Getenv("STUB") != "" {
		stub = true
	}

	Constants = constants{
		ServiceName: "bbqberry",
		Version:     "v1",
		SteinhartHart: steinhartHart{
			// From http://www.thinksrs.com/downloads/programs/Therm%20Calc/NTCCalibrator/NTCcalculator.htm
			// A=64, V=0.20625, R=15000, K=0.00036978864, C=-277.14963, F=-466.86932
			// A: 0.9114243730,
			// B: 1.915917966,
			// C: 1.445400011,
			//
			// From http://www.thermistor.com/calculators.php
			// Z/D (-4.4%/C @ 25C) Mil Ratio B 	99.6K	(iGrill 2 Probe)
			A: 0.000535683116684,
			B: 0.000240497428160, //(DOES NOT WORK)
			C: 0.000000039680880,
			//
			// From http://www.thermistor.com/calculators.php
			// Z/D (-4.4%/C @ 25C) Mil Ratio B 	968K	(PR-002 Ambient)
			// A=78, V=0.25136718, R=12128, K=446.0333, C=168.8833, F=335.98993		(Actual ~330)
			// A: 0.000001913640634,
			// B: 0.000238220028939,					//(WORKS)
			// C: 0.000000031739980,
		},
		Stub: stub,
	}
}

// SteinhartHart holds the co-efficients for the SteinhartHart temperature calculations
type steinhartHart struct {
	A, B, C float64
}

type constants struct {
	ServiceName   string
	Version       string
	SteinhartHart steinhartHart
	Stub          bool
}

// Constants contains static information about the running service
var Constants constants
