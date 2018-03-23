package stubembd

import "github.com/Polarishq/middleware/framework/log"

// StubSPIBus provides a stub of embd.SPIBus
type StubSPIBus struct {
	simulateTemp bool
	CloseCallCount, ReceiveByteCallCount, ReceiveDataCallCount,
	TransferAndReceiveByteCallCount, TransferAndReceiveDataCallCount, WriteCallCount int
}

// NewStubSPIBus creates a new stubbed SPIBus
func NewStubSPIBus(channel byte) *StubSPIBus {
	return &StubSPIBus{
		simulateTemp: channel == 0,
	}
}

// ResetCallCounts resets the call counts of all embd.SPIBus interface methods
func (o *StubSPIBus) Reset() {
	o.CloseCallCount = 0
	o.WriteCallCount = 0
	o.ReceiveByteCallCount = 0
	o.ReceiveDataCallCount = 0
	o.TransferAndReceiveByteCallCount = 0
	o.TransferAndReceiveDataCallCount = 0

	resetFakeTemps()
}

// Close - See embd.SPIBus
func (o *StubSPIBus) Close() error {
	log.Info("Close")

	o.CloseCallCount++
	return nil
}

// ReceiveByte - See embd.SPIBus
func (o *StubSPIBus) ReceiveByte() (byte, error) {
	log.Info("ReceiveByte")

	o.ReceiveByteCallCount++
	return 1, nil
}

// ReceiveData - See embd.SPIBus
func (o *StubSPIBus) ReceiveData(data int) ([]byte, error) {
	log.WithFields(log.Fields{
		"data": data,
	}).Info("ReceiveData")

	o.ReceiveDataCallCount++
	return nil, nil
}

// TransferAndReceiveByte - See embd.SPIBus
func (o *StubSPIBus) TransferAndReceiveByte(data byte) (byte, error) {
	log.WithFields(log.Fields{
		"data": data,
	}).Info("TransferAndReceiveByte")

	o.TransferAndReceiveByteCallCount++
	return 1, nil
}

// TransferAndReceiveData - See embd.SPIBus
func (o *StubSPIBus) TransferAndReceiveData(data []byte) error {
	log.WithFields(log.Fields{
		"data": data,
	}).Info("TransferAndReceiveData")

	if o.simulateTemp {
		a := getFakeTemp(0)
		data[0] = byte(a & 0x00FF)			// low byte
		data[1] = byte(a >> 8 & 0x03)		// high byte
		data[0] = 0x0
	}

	o.TransferAndReceiveDataCallCount++
	return nil
}

// Write - See embd.SPIBus
func (o *StubSPIBus) Write(data []byte) (int, error) {
	o.WriteCallCount++
	log.WithFields(log.Fields{
		"data": data,
	}).Info("Write")
	return 1, nil
}
