package mocks

import (
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

// SatNumRepositoryMock คือ Mock struct ที่ใช้แทน ports.SatNumRepositoryPort
type SatNumRepositoryMock struct {
	mock.Mock
}

// GetAllSatNums จำลองการดึงข้อมูลทั้งหมด
func (m *SatNumRepositoryMock) GetAllSatNums() ([]domain.SatNum, error) {
	// Call เป็นฟังก์ชันของ testify/mock ที่จะบันทึกว่า method นี้ถูกเรียก
	args := m.Called()

	// args.Get(0) จะดึงค่าที่คาดหวัง (Mock Value) ที่เราตั้งไว้ในการทดสอบ
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.SatNum), args.Error(1)
}

// GetSatNumByCharKey จำลองการดึงข้อมูลตาม CharKey
func (m *SatNumRepositoryMock) GetSatNumByCharKey(key string) (*domain.SatNum, error) {
	args := m.Called(key)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.SatNum), args.Error(1)
}
