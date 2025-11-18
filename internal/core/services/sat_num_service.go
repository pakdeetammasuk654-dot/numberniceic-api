package services

import (
	"errors"

	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/domain"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/ports"
)

// satNumService คือโครงสร้างที่ใช้ Implement ports.SatNumServicePort
type satNumService struct {
	// repo คือ Dependency ที่เป็น Repository Port
	// ทำให้ Service Layer ไม่ต้องรู้ว่า Repository ตัวจริงคืออะไร (PostgreSQL, Mock, ฯลฯ)
	repo ports.SatNumRepositoryPort
}

// NewSatNumService เป็น Constructor สำหรับสร้าง Service ใหม่
// โดยรับ Repository Port เข้ามาผ่าน Dependency Injection
func NewSatNumService(repo ports.SatNumRepositoryPort) ports.SatNumServicePort {
	return &satNumService{repo: repo}
}

// FetchAll Implement method จาก SatNumServicePort
// เรียกใช้ Repository เพื่อดึงข้อมูลทั้งหมด
func (s *satNumService) FetchAll() ([]domain.SatNum, error) {
	// ใน Core Service เราสามารถเพิ่ม Business Logic เช่น การตรวจสอบสิทธิ์, การ Cache, หรือการกรองข้อมูลก่อนส่งออก

	// ในตัวอย่างนี้ เราแค่เรียก Repository Port ตรงๆ
	satNums, err := s.repo.GetAllSatNums()
	if err != nil {
		return nil, errors.New("failed to fetch all sat nums from data source")
	}

	return satNums, nil
}

// GetByKey Implement method จาก SatNumServicePort
func (s *satNumService) GetByKey(key string) (*domain.SatNum, error) {
	if key == "" {
		return nil, errors.New("char key cannot be empty")
	}

	satNum, err := s.repo.GetSatNumByCharKey(key)
	if err != nil {
		// ที่นี่เราอาจจะตรวจสอบว่าเป็น Error แบบ "Not Found" หรือ "Internal Error"
		return nil, errors.New("sat num not found or data source error")
	}

	return satNum, nil
}
