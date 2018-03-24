package ads1x15

import (
	"github.com/kidoman/embd"
	"time"
	"fmt"
	"github.com/Polarishq/middleware/framework/log"
)

type ads1x15 struct {
	bus    	embd.I2CBus
	addr 	byte
}

// NewADS1115 creates a new object capable of communicating with a NewADS1115 chip over I2C
func NewADS1115(addr byte, bus embd.I2CBus) *ads1x15 {
	return &ads1x15{
		bus: bus,
		addr: addr,
	}
}

/*
    def _read(self, mux, gain, data_rate, mode):
        """Perform an ADC read with the provided mux, gain, data_rate, and mode
        values.  Returns the signed integer result of the read.
        """
        config = ADS1x15_CONFIG_OS_SINGLE  # Go out of power-down mode for conversion.
        # Specify mux value.
        config |= (mux & 0x07) << ADS1x15_CONFIG_MUX_OFFSET
        # Validate the passed in gain and then set it in the config.
        if gain not in ADS1x15_CONFIG_GAIN:
            raise ValueError('Gain must be one of: 2/3, 1, 2, 4, 8, 16')
        config |= ADS1x15_CONFIG_GAIN[gain]
        # Set the mode (continuous or single shot).
        config |= mode
        # Get the default data rate if none is specified (default differs between
        # ADS1015 and ADS1115).
        if data_rate is None:
            data_rate = self._data_rate_default()
        # Set the data rate (this is controlled by the subclass as it differs
        # between ADS1015 and ADS1115).
        config |= self._data_rate_config(data_rate)
        config |= ADS1x15_CONFIG_COMP_QUE_DISABLE  # Disble comparator mode.
        # Send the config value to start the ADC conversion.
        # Explicitly break the 16-bit value down to a big endian pair of bytes.
        self._device.writeList(ADS1x15_POINTER_CONFIG, [(config >> 8) & 0xFF, config & 0xFF])
        # Wait for the ADC sample to finish based on the sample rate plus a
        # small offset to be sure (0.1 millisecond).
        time.sleep(1.0/data_rate+0.0001)
        # Retrieve the result.
        result = self._device.readList(ADS1x15_POINTER_CONVERSION, 2)
        return self._conversion_value(result[1], result[0])

 */
func (o *ads1x15) read(mux, gain, mode int) (int, error) {
	config := ADS1x15_CONFIG_OS_SINGLE  // Go out of power-down mode for conversion.

	// Specify mux value.
	config |= (mux & 0x07) << ADS1x15_CONFIG_MUX_OFFSET

	// Validate the passed in gain and then set it in the config.
	gainv, ok := ADS1x15_CONFIG_GAIN[gain]
	if ! ok {
		return ADS1x15_READ_FAIL, fmt.Errorf("Invalid gain %d, must be one of %v", gain, ADS1x15_CONFIG_GAIN)
	}
	config |= ADS1x15_CONFIG_GAIN[gainv]

	// Set the mode (continuous or single shot).
	config |= mode

	// Set the data rate (this is controlled by the subclass as it differs between ADS1015 and ADS1115).
	data_rate := ADS1115_CONFIG_DR_DEFAULT
	config |= ADS1115_CONFIG_DR[data_rate]

	config |= ADS1x15_CONFIG_COMP_QUE_DISABLE  // Disble comparator mode.

	// Send the config value to start the ADC conversion.
	// Explicitly break the 16-bit value down to a big endian pair of bytes.
	//self._device.writeList(ADS1x15_POINTER_CONFIG, [(config >> 8) & 0xFF, config & 0xFF])
	err := o.bus.WriteToReg(o.addr, ADS1x15_POINTER_CONFIG, []byte{byte(config >> 8) & 0xFF, byte(config & 0xFF)})
	if err != nil {
		return ADS1x15_READ_FAIL, err
	}

	// Wait for the ADC sample to finish based on the sample rate plus a
	// small offset to be sure (0.1 millisecond).
	// TODO: Check the value of this, i.e: nano seconds vs milliseconds
	time.Sleep(time.Duration(1.0 / data_rate + int(time.Millisecond)))

	// Retrieve the result.
	//result = self._device.readList(ADS1x15_POINTER_CONVERSION, 2)
	result := make([]byte, 2)
	err = o.bus.ReadFromReg(o.addr, ADS1x15_POINTER_CONVERSION, result)
	if err != nil {
		return ADS1x15_READ_FAIL, err
	}

	return o.convert(result[1], result[0]), nil
}

/*
    def _conversion_value(self, low, high):
        # Convert to 16-bit signed value.
        value = ((high & 0xFF) << 8) | (low & 0xFF)
        # Check for sign bit and turn into a negative value if set.
        if value & 0x8000 != 0:
            value -= 1 << 16
        return value

 */
func (o *ads1x15) convert(low, high byte) int {
	value := int(high) & 0xFF << 8 | int(low & 0xFF)

	log.WithFields(log.Fields{
		"low": fmt.Sprintf("%02x", low),
		"high": fmt.Sprintf("%02x", high),
		"value": value,
	}).Info("convert")

	if value & 0x8000 != 0 {
		value -= 1 << 16
		log.WithFields(log.Fields{
			"value": value,
		}).Info("value munged - figure out why this is here")
	}

	return value
}

func (o *ads1x15) AnalogValueAt(chanNum int) (int, error) {
	result, err := o.read(int(chanNum + 0x04), 1, ADS1x15_CONFIG_MODE_SINGLE)
	if err != nil {
		return ADS1x15_READ_FAIL, err
	}

	return result, nil
}

