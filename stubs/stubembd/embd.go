package stubembd

// StubSPIBus provides a stub of embd.SPIBus
type StubSPIBus struct {
	CloseCallCount, ReceiveByteCallCount, ReceiveDataCallCount, 
	TransferAndReceiveByteCallCount, TransferAndReceiveDataCallCount, WriteCallCount int		
}

// ResetCallCounts resets the call counts of all embd.SPIBus interface methods
func (_m *StubSPIBus) ResetCallCounts() {
	_m.CloseCallCount = 0
	_m.ReceiveByteCallCount = 0
	_m.ReceiveDataCallCount = 0
	_m.TransferAndReceiveByteCallCount = 0
	_m.TransferAndReceiveDataCallCount = 0
}

// NewStubSPIBus creates a new stubbed SPIBus
func NewStubSPIBus() *StubSPIBus {
	return &StubSPIBus{}
}

// Close - See embd.SPIBus
func (_m *StubSPIBus) Close() error {
	_m.CloseCallCount++
	return nil
}

// ReceiveByte - See embd.SPIBus
func (_m *StubSPIBus) ReceiveByte() (byte, error) {
	_m.ReceiveByteCallCount++
	return 1, nil
}

// ReceiveData - See embd.SPIBus
func (_m *StubSPIBus) ReceiveData(_param0 int) ([]byte, error) {
	_m.ReceiveDataCallCount++
	return nil, nil
}

// TransferAndReceiveByte - See embd.SPIBus
func (_m *StubSPIBus) TransferAndReceiveByte(_param0 byte) (byte, error) {
	_m.TransferAndReceiveByteCallCount++
	return 1, nil
}

// TransferAndReceiveData - See embd.SPIBus
func (_m *StubSPIBus) TransferAndReceiveData(_param0 []byte) error {
	_m.TransferAndReceiveDataCallCount++
	return nil
}

// Write - See embd.SPIBus
func (_m *StubSPIBus) Write(_param0 []byte) (int, error) {
	_m.WriteCallCount++
	return 1, nil
}
