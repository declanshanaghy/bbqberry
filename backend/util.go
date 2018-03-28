package backend

import "github.com/declanshanaghy/bbqberry/framework"

// getProbeIndexes returns a slice containing all the probe numbers to query, based on the given probe parameter
// which should one of:
// a single probe index: query a single probe
// nil or < 0: query all probes
func getProbeIndexes(n *int32) (probes []int32) {
	l := len(framework.Config.Hardware.Probes)
	if n == nil || *n < 0 {
		probes = make([]int32, l)
		for i := 0; i < l; i++ {
			probes[i] = int32(i)
		}
	} else {
		probes = make([]int32, 1)
		probes[0] = *n
	}
	return
}
