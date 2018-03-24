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
	log.WithField("channel", channel).Info("NewStubSPIBus")
	return &StubSPIBus{
		simulateTemp: channel == 1,
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
	//log.WithFields(log.Fields{
	//	"simulateTemp": o.simulateTemp,
	//	"data": data,
	//}).Debugf("TransferAndReceiveData - Transfer")

	if o.simulateTemp {
		a := getFakeTemp(0)
		data[0] = byte(a & 0x00FF)			// low byte
		data[1] = byte(a >> 8 & 0x03)		// high byte
	}

	//log.WithFields(log.Fields{
	//	"data": data,
	//}).Debugf("TransferAndReceiveData - Receive")

	o.TransferAndReceiveDataCallCount++
	return nil
}

// Write - See embd.SPIBus
func (o *StubSPIBus) Write(data []byte) (int, error) {
	o.WriteCallCount++
	//log.WithFields(log.Fields{
	//	"data": data,
	//}).Debugf("Write")
	return 1, nil
}
