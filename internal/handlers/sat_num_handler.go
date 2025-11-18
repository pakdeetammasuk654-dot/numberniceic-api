package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/ports"
)

// SatNumHandler คือ Adapter ที่จัดการ HTTP Request/Response
type SatNumHandler struct {
	// service คือ Input Port ของ Core Service
	service ports.SatNumServicePort
}

// NewSatNumHandler คือ Constructor สำหรับ Handler
func NewSatNumHandler(service ports.SatNumServicePort) *SatNumHandler {
	return &SatNumHandler{service: service}
}

// GetAllSatNums จัดการ GET /api/v1/sat-nums (ดึงข้อมูลทั้งหมด)
func (h *SatNumHandler) GetAllSatNums(c *fiber.Ctx) error {
	// 1. เรียกใช้ Core Service ผ่าน Port
	satNums, err := h.service.FetchAll()
	if err != nil {
		// ส่ง 500 Internal Server Error หาก Service เกิดข้อผิดพลาด
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error: Could not fetch data",
		})
	}

	// 2. ส่ง 200 OK พร้อมข้อมูล
	return c.Status(http.StatusOK).JSON(satNums)
}

// GetSatNumByCharKey จัดการ GET /api/v1/sat-nums/:key (ดึงข้อมูลตามคีย์)
func (h *SatNumHandler) GetSatNumByCharKey(c *fiber.Ctx) error {
	// 1. อ่าน CharKey จาก URL parameter
	key := c.Params("key")
	if key == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'key' parameter",
		})
	}

	// 2. เรียกใช้ Core Service ผ่าน Port
	satNum, err := h.service.GetByKey(key)

	if err != nil {
		// จาก SatNum Service เรากำหนดให้ return Error เมื่อไม่พบข้อมูล
		// (ใน Service Layer จริงๆ ควรมีการตรวจสอบ Error Type ที่ละเอียดกว่านี้)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "SatNum not found or invalid key format",
		})
	}

	// 3. ส่ง 200 OK พร้อมข้อมูล
	return c.Status(http.StatusOK).JSON(satNum)
}
