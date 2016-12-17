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
			// From http://www.thermistor.com/calculators.php
			// 
			// iDevices iGrill2 Probe
			// Lower		-50
			// Upper		150
			// Resistance	101466	(Measured in circuit)
			// // Curve		R (-6.2%/C @ 25C) Mil Ratio X
			// A: 0.001491353519147,
			// B: 0.000152056148089,
			// C: 0.000000071310824,
			//
			// Lower		-50
			// Upper		250
			// Resistance	101466	(Measured in circuit)
			// Curve		Z/D (-4.4%/C @ 25C) Mil Ratio B
			A: 0.000589420497753,
			B: 0.000231533027075,
			C: 0.000000063975391,
			//
			// Home Circles Grill Probe
			// Lower				-50
			// Upper				150
			// Resistance at 25c	231257.17188	(Measured in circuit)
			// Curve				R (-6.2%/C @ 25C) Mil Ratio X
			// A: 0.001372385462693,
			// B: 0.000150289967635,
			// C: 0.000000066337977,
			//
			// Lower				-50
			// Upper				150
			// Resistance at 25c	231257.17188	(Measured in circuit)
			// Curve				R (-6.2%/C @ 25C) Mil Ratio X
			// A: 0.000419836753426,
			// B: 0.000227935261483,
			// C: 0.000000063356527,
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
	Rn      float32
}

type constants struct {
	ServiceName   string
	Version       string
	SteinhartHart steinhartHart
	Stub          bool
}

// Constants contains static information about the running service
var Constants constants
