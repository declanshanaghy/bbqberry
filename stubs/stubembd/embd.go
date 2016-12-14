package stubembd

// StubSPIBus provides a stub of embd.SPIBus
type StubSPIBus struct {
}

// NewStubSPIBus creates a new stubbed SPIBus
func NewStubSPIBus() *StubSPIBus {
	return &StubSPIBus{}
}

// Close - See embd.SPIBus
func (_m *StubSPIBus) Close() error {
	return nil
}

// ReceiveByte - See embd.SPIBus
func (_m *StubSPIBus) ReceiveByte() (byte, error) {
	return 1, nil
}

// ReceiveData - See embd.SPIBus
func (_m *StubSPIBus) ReceiveData(_param0 int) ([]byte, error) {
	return nil, nil
}

// TransferAndReceiveByte - See embd.SPIBus
func (_m *StubSPIBus) TransferAndReceiveByte(_param0 byte) (byte, error) {
	return 1, nil
}

// TransferAndReceiveData - See embd.SPIBus
func (_m *StubSPIBus) TransferAndReceiveData(_param0 []byte) error {
	return nil
}

// Write - See embd.SPIBus
func (_m *StubSPIBus) Write(_param0 []byte) (int, error) {
	return 1, nil
}
