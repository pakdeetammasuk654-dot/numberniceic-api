package services_test

import (
	"errors"
	"testing"

	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/domain"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/services"
	mocks "github.com/pakdeetammasuk654-dot/numberniceic-api/internal/testify/mock"
	"github.com/stretchr/testify/assert"
)

// TestFetchAll_Success ทดสอบกรณีดึงข้อมูลทั้งหมดสำเร็จ
func TestFetchAll_Success(t *testing.T) {
	// 1. Setup Mock Data
	mockSatNums := []domain.SatNum{
		{CharKey: "A", SatValue: 1},
		{CharKey: "B", SatValue: 2},
	}

	// 2. Initialize Mock Repository
	mockRepo := new(mocks.SatNumRepositoryMock)

	// 3. Define Mock Behavior (เมื่อ GetAllSatNums ถูกเรียก ให้ return mockSatNums และ nil error)
	mockRepo.On("GetAllSatNums").Return(mockSatNums, nil).Once()

	// 4. Initialize Service (Inject Mock Repo)
	satNumService := services.NewSatNumService(mockRepo)

	// 5. Execute Method
	result, err := satNumService.FetchAll()

	// 6. Assertions (ตรวจสอบผลลัพธ์)
	assert.Nil(t, err)                   // ต้องไม่มี error
	assert.Equal(t, mockSatNums, result) // ผลลัพธ์ต้องตรงกับ mock data
	mockRepo.AssertExpectations(t)       // ตรวจสอบว่า mock method ถูกเรียกจริงตามที่ตั้งค่า
}

// TestGetByKey_Success ทดสอบกรณีดึงข้อมูลตาม Key สำเร็จ
func TestGetByKey_Success(t *testing.T) {
	mockKey := "C"
	mockSatNum := domain.SatNum{CharKey: mockKey, SatValue: 3}

	mockRepo := new(mocks.SatNumRepositoryMock)

	// กำหนดให้เมื่อ GetSatNumByCharKey ถูกเรียกด้วย "C" ให้ return mockSatNum และ nil error
	mockRepo.On("GetSatNumByCharKey", mockKey).Return(&mockSatNum, nil).Once()

	satNumService := services.NewSatNumService(mockRepo)

	result, err := satNumService.GetByKey(mockKey)

	assert.Nil(t, err)
	assert.Equal(t, &mockSatNum, result)
	mockRepo.AssertExpectations(t)
}

// TestGetByKey_EmptyKey ทดสอบกรณีส่ง CharKey ว่าง
func TestGetByKey_EmptyKey(t *testing.T) {
	mockRepo := new(mocks.SatNumRepositoryMock)
	satNumService := services.NewSatNumService(mockRepo)

	// ไม่ต้องกำหนด Mock Behavior เพราะ service ควรจะตรวจสอบและ return error ก่อนถึง Repository

	result, err := satNumService.GetByKey("")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "char key cannot be empty") // ตรวจสอบข้อความ error
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetSatNumByCharKey") // ตรวจสอบว่า Repository ไม่ถูกเรียก
}

// TestGetByKey_RepoError ทดสอบกรณี Repository เกิด Error
func TestGetByKey_RepoError(t *testing.T) {
	mockKey := "D"
	mockError := errors.New("database connection failed")

	mockRepo := new(mocks.SatNumRepositoryMock)

	// กำหนดให้เมื่อถูกเรียก ให้ return nil domain และ mockError
	mockRepo.On("GetSatNumByCharKey", mockKey).Return(nil, mockError).Once()

	satNumService := services.NewSatNumService(mockRepo)

	result, err := satNumService.GetByKey(mockKey)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "sat num not found or data source error") // ตรวจสอบว่า Service แปลง Error ได้ถูกต้อง
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
