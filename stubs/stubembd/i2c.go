package stubembd

import (
	"github.com/Polarishq/middleware/framework/log"
	"fmt"
)

// StubI2CBus provides a stub of embd.I2CBus
type StubI2CBus struct {
	CloseCallCount, WriteToRegCallCount, ReadFromRegCallCount int
	ReadFromRegReturnValues [][]int
}

// NewStubI2CBus creates a new stubbed SPIBus
func NewStubI2CBus() *StubI2CBus {
	return &StubI2CBus{}
}

// ResetCallCounts resets the call counts of all embd.SPIBus interface methods
func (o *StubI2CBus) Reset() {
	o.CloseCallCount = 0
	o.WriteToRegCallCount = 0
	o.ReadFromRegCallCount = 0

	resetFakeTemps()
}

// ReadByte reads a byte from the given address.
func (o *StubI2CBus) ReadByte(addr byte) (value byte, err error) {
	return 0,nil
}

// ReadBytes reads a slice of bytes from the given address.
func (o *StubI2CBus) ReadBytes(addr byte, num int) (value []byte, err error) {
	return nil,nil
}

// WriteByte writes a byte to the given address.
func (o *StubI2CBus) WriteByte(addr, value byte) error {
	return nil
}

// WriteBytes writes a slice bytes to the given address.
func (o *StubI2CBus) WriteBytes(addr byte, value []byte) error {
	return nil
}

// ReadFromReg reads n (len(value)) bytes from the given address and register.
func (o *StubI2CBus) ReadFromReg(addr, reg byte, value []byte) error {
	//Assume probe 0 here
	a := getFakeTemp(0)

	o.ReadFromRegCallCount++

	value[1] = byte(a & 0x00FF)			// low byte
	value[0] = byte((a >> 8) & 0x00FF)	// high byte

	log.WithFields(log.Fields{
		"reg": reg,
		"a": a,
		"value1": fmt.Sprintf("%02x", value[1]),
		"value0": fmt.Sprintf("%02x", value[0]),
	}).Debugf("ReadFromReg")

	return nil
}

// ReadByteFromReg reads a byte from the given address and register.
func (o *StubI2CBus) ReadByteFromReg(addr, reg byte) (value byte, err error) {
	return 0,nil
}

// ReadU16FromReg reads a unsigned 16 bit integer from the given address and register.
func (o *StubI2CBus) ReadWordFromReg(addr, reg byte) (value uint16, err error) {
	return 0,nil
}

// WriteToReg writes len(value) bytes to the given address and register.
func (o *StubI2CBus) WriteToReg(addr, reg byte, value []byte) error {
	o.WriteToRegCallCount++
	return nil
}

// WriteByteToReg writes a byte to the given address and register.
func (o *StubI2CBus) WriteByteToReg(addr, reg, value byte) error {
	return nil
}

// WriteU16ToReg
func (o *StubI2CBus) WriteWordToReg(addr, reg byte, value uint16) error {
	return nil
}

// Close releases the resources associated with the bus.
func (o *StubI2CBus) Close() error {
	o.CloseCallCount++
	return nil
}
