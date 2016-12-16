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
			// iDevices iGrill2 Probe
			// Lower		-50
			// Upper		150
			// Resistance	102700
			// // Curve		R (-6.2%/C @ 25C) Mil Ratio X	
			A: 0.001489592596173,
			B: 0.000152030416280,
			C: 0.000000071232470,
			//
			// Lower		-50
			// Upper		150
			// Resistance	102700
			// // Curve		Z/D (-4.4%/C @ 25C) Mil Ratio B	
			// A: 0.000599183607286,
			// B: 0.000229721026267,
			// C: 0.000000067919036,
			// 
			// From http://www.thermistor.com/calculators.php
			// Home Circles Grill Probe
			// Lower				-50
			// Upper				150
			// Resistance at 25c	197000
			// Curve				Z/D (-4.4%/C @ 25C) Mil Ratio B
			// A: 0.000455101186551,
			// B: 0.000228288824313,
			// C: 0.000000064208459,
			//			
			// Lower				-50
			// Upper				150
			// Resistance at 25c	204000
			// Curve				R (-6.2%/C @ 25C) Mil Ratio X
			// A: 0.001390360423839,
			// B: 0.000150560390246,
			// C: 0.000000067049772,			
		},
		Stub: stub,
	}
}

// SteinhartHart holds the co-efficients for the SteinhartHart temperature calculations
type steinhartHart struct {
	A, B, C float64
	Rn float32
}

type constants struct {
	ServiceName   string
	Version       string
	SteinhartHart steinhartHart
	Stub          bool
}

// Constants contains static information about the running service
var Constants constants
