package stubembd

// StubSPIBus provides a stub of embd.SPIBus
type StubSPIBus struct {
	CloseCallCount, ReceiveByteCallCount, ReceiveDataCallCount,
	TransferAndReceiveByteCallCount, TransferAndReceiveDataCallCount, WriteCallCount int
}

// ResetCallCounts resets the call counts of all embd.SPIBus interface methods
func (o *StubSPIBus) ResetCallCounts() {
	o.CloseCallCount = 0
	o.ReceiveByteCallCount = 0
	o.ReceiveDataCallCount = 0
	o.TransferAndReceiveByteCallCount = 0
	o.TransferAndReceiveDataCallCount = 0
}

// NewStubSPIBus creates a new stubbed SPIBus
func NewStubSPIBus() *StubSPIBus {
	return &StubSPIBus{}
}

// Close - See embd.SPIBus
func (o *StubSPIBus) Close() error {
	o.CloseCallCount++
	return nil
}

// ReceiveByte - See embd.SPIBus
func (o *StubSPIBus) ReceiveByte() (byte, error) {
	o.ReceiveByteCallCount++
	return 1, nil
}

// ReceiveData - See embd.SPIBus
func (o *StubSPIBus) ReceiveData(_param0 int) ([]byte, error) {
	o.ReceiveDataCallCount++
	return nil, nil
}

// TransferAndReceiveByte - See embd.SPIBus
func (o *StubSPIBus) TransferAndReceiveByte(_param0 byte) (byte, error) {
	o.TransferAndReceiveByteCallCount++
	return 1, nil
}

// TransferAndReceiveData - See embd.SPIBus
func (o *StubSPIBus) TransferAndReceiveData(_param0 []byte) error {
	o.TransferAndReceiveDataCallCount++
	return nil
}

// Write - See embd.SPIBus
func (o *StubSPIBus) Write(_param0 []byte) (int, error) {
	o.WriteCallCount++
	return 1, nil
}
