package ports

import "github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/domain"

// SatNumRepositoryPort คือ Output Port สำหรับ Repository/Database Adapter
// กำหนด method ที่ Core Service ต้องการใช้ในการเข้าถึงข้อมูล
type SatNumRepositoryPort interface {
	// GetAllSatNums ดึงข้อมูลทั้งหมดจากตาราง sat_nums
	GetAllSatNums() ([]domain.SatNum, error)

	// GetSatNumByCharKey ดึงข้อมูล SatNum จาก CharKey ที่กำหนด
	GetSatNumByCharKey(key string) (*domain.SatNum, error)
}

// SatNumServicePort คือ Input Port สำหรับ Handler/Service
// กำหนด method ที่ HTTP Handler จะเรียกใช้ Business Logic
type SatNumServicePort interface {
	// FetchAll ดึงข้อมูล SatNum ทั้งหมดเพื่อนำไปประมวลผลหรือแสดงผล
	FetchAll() ([]domain.SatNum, error)

	// GetByKey ดึงข้อมูล SatNum โดยใช้ CharKey
	GetByKey(key string) (*domain.SatNum, error)
}
