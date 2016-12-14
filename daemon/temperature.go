package daemon

import "time"
import "github.com/declanshanaghy/bbqberry/hardware"
import "github.com/declanshanaghy/bbqberry/models"
import "github.com/golang/glog"

// CollectAndLogTermperatureMetrics executes a continuous loop reading temperature sensors and logging to InfluxDB
func CollectAndLogTermperatureMetrics(quit <-chan bool) {
    glog.Info("action=start")
    t := time.NewTicker(time.Second)
    temp := hardware.NewTemperatureReader()

    for true {
        select {
            case <-quit:
                glog.Info("action=quit")
                break;
            case <-t.C:
                glog.V(2).Infoln("action=timeout")
                readings := collectTermperatureMetrics(temp)
                logTermperatureMetrics(readings)
        }
    }

    t.Stop()
    temp.Close()
    glog.Info("action=done")
}

func collectTermperatureMetrics(temp hardware.TemperatureArray) *models.TemperatureReadings {
    glog.V(2).Infof("action=start numProbes=%d", temp.GetNumProbes())
    readings := models.TemperatureReadings{}
    for i := int32(1); i <= temp.GetNumProbes(); i++ {
        glog.V(3).Infof("action=iterate probe=%d", i)
        reading := models.TemperatureReading{}
        if err := temp.GetTemperatureReading(i, &reading); err != nil {
            glog.Error(err)
        }
        readings = append(readings, &reading)
    }
    glog.V(2).Infof("action=done")
    return &readings
}

func logTermperatureMetrics(readings *models.TemperatureReadings) {
    glog.V(2).Infof("action=start numReadings=%d", len(*readings))
    glog.V(2).Infof("action=done")
}