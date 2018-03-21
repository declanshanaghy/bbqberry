package hardware

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"time"
)

/*
def write(self, data, assert_ss=True, deassert_ss=True):
	"""Half-duplex SPI write.  If assert_ss is True, the SS line will be
	asserted low, the specified bytes will be clocked out the MOSI line, and
	if deassert_ss is True the SS line be put back high.
	"""
	# Fail MOSI is not specified.
	if self._mosi is None:
		raise RuntimeError('Write attempted with no MOSI pin specified.')
	if assert_ss and self._ss is not None:
		self._gpio.set_low(self._ss)
	for byte in data:
		for i in range(8):
			# Write bit to MOSI.
			if self._write_shift(byte, i) & self._mask:
				self._gpio.set_high(self._mosi)
			else:
				self._gpio.set_low(self._mosi)
			# Flip clock off base.
			self._gpio.output(self._sclk, not self._clock_base)
			# Return clock to base.
			self._gpio.output(self._sclk, self._clock_base)
	if deassert_ss and self._ss is not None:
		self._gpio.set_high(self._ss)
 */

type SoftwareSPI struct {
	clk 		int
	mosi 	int
}

func NewSoftwareSPI(clk, mosi int) *SoftwareSPI {
	o := SoftwareSPI{
		clk:	clk,
		mosi:	mosi,

	}
	embd.SetDirection(o.clk, embd.Out)
	embd.SetDirection(o.mosi, embd.Out)

	return &o
}

// TransferAndReceiveData transmits data in a buffer(slice) and receives into it.
func (o *SoftwareSPI) TransferAndReceiveData(dataBuffer []uint8) error {
	panic("Not implemented")
}

// ReceiveData receives data of length len into a slice.
func (o *SoftwareSPI) ReceiveData(len int) ([]uint8, error) {
	panic("Not implemented")
}

// TransferAndReceiveByte transmits a byte data and receives a byte.
func (o *SoftwareSPI) TransferAndReceiveByte(data byte) (byte, error) {
	panic("Not implemented")
}

// ReceiveByte receives a byte data.
func (o *SoftwareSPI) ReceiveByte() (byte, error) {
	panic("Not implemented")
}

// Close releases the resources associated with the bus.
func (o *SoftwareSPI) Close() error {
	return nil
}

func (o *SoftwareSPI) Write(data []byte) (n int, err error) {
	// Start of transfer
	// Ensure clock is low
	embd.DigitalWrite(o.clk, embd.Low)
	time.Sleep(time.Millisecond)

	// Bring chip select low
	//embd.DigitalWrite(o.cs, embd.Low)

	//log.Debugf("Sending %d bytes", len(data))

	for _, by := range(data) {
		//log.Debugf("Sending %02x, %08b", by, by)
		for bi := uint(0); bi < 8; bi ++ {
			m := uint(1) << bi
			b := (uint(by) & m) == m
			//log.Debugf("Bit %d of %08b is %b", bi, by, b)
			if b {
				embd.DigitalWrite(o.mosi, embd.High)
			}  else {
				embd.DigitalWrite(o.mosi, embd.Low)
			}

			embd.DigitalWrite(o.clk, embd.High)
			//time.Sleep(time.Nanosecond * 100)
			embd.DigitalWrite(o.clk, embd.Low)
		}
	}

	// End of transfer
	// Ensure clock is low
	embd.DigitalWrite(o.clk, embd.High)
	// Bring chip select high
	//embd.DigitalWrite(o.cs, embd.High)

	//log.Debugf("Done sending %d bytes", len(data))

	return len(data), nil
}