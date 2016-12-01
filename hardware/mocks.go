package hardware

import (
	"github.com/Polarishq/middleware/framework/log"
)

// SPIBus interface allows interaction with the SPI bus.
type MockSPIBus struct {}

// TransferAndReceiveData transmits data in a buffer(slice) and receives into it.
func (b *MockSPIBus) TransferAndReceiveData(dataBuffer []uint8) error {
	log.Infof("action=TransferAndReceiveData dataBuffer=%v", dataBuffer)
	return nil
}
	
// ReceiveData receives data of length len into a slice.
func (b *MockSPIBus) ReceiveData(len int) ([]uint8, error) {
	log.Infof("action=ReceiveData len=%d", len)
	return make([]uint8, len), nil
}
	
// TransferAndReceiveByte transmits a byte data and receives a byte.
func (b *MockSPIBus) TransferAndReceiveByte(data byte) (byte, error) {
	log.Infof("action=TransferAndReceiveByte data=%v", data)
	return 0, nil
}
	
// ReceiveByte receives a byte data.
func (b *MockSPIBus) ReceiveByte() (byte, error) {
	log.Info("action=ReceiveByte")
	return 0, nil
	
}
	
// Close releases the resources associated with the bus.
func (b *MockSPIBus) Close() error {
	log.Info("action=Close")
	return nil
}

func (b *MockSPIBus) Write(p []byte) (int, error) {
	log.Info("action=Write p=%v", p)
	return 0, nil
}
