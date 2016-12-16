package daemon

import "time"
import "github.com/declanshanaghy/bbqberry/hardware"
import "github.com/declanshanaghy/bbqberry/models"
import "github.com/Polarishq/middleware/framework/log"
import "sync"

// CollectAndLogTermperatureMetrics executes a continuous loop reading temperature sensors and logging to InfluxDB
func CollectAndLogTermperatureMetrics(run <-chan bool, wg *sync.WaitGroup) {
	log.Info("action=start")
	defer wg.Done()

	t := time.NewTicker(time.Second * 10)
	temp := hardware.NewTemperatureReader()
	loop := true

	// Log immediately then enter the loop
	collectTemperatureMetrics(temp)

	for loop {
		select {
		case loop = <-run:
			log.Infof("action=rx loop=%t", loop)
		case <-t.C:
			log.Debugf("action=timeout")
			readings := collectTemperatureMetrics(temp)
			logTemperatureMetrics(readings)
		}
	}

	t.Stop()
	temp.Close()
	log.Info("action=done")
}

func collectTemperatureMetrics(temp hardware.TemperatureArray) *models.TemperatureReadings {
	log.Infof("action=start numProbes=%d", temp.GetNumProbes())
	readings := models.TemperatureReadings{}
	for i := int32(1); i <= temp.GetNumProbes(); i++ {
		log.Debugf("action=iterate probe=%d", i)
		reading := models.TemperatureReading{}
		if err := temp.GetTemperatureReading(i, &reading); err != nil {
			log.Error(err)
		}
		readings = append(readings, &reading)
	}
	log.Infof("action=done")
	return &readings
}

func logTemperatureMetrics(readings *models.TemperatureReadings) {
	log.Infof("action=start numReadings=%d", len(*readings))
	log.Infof("action=done")
}
