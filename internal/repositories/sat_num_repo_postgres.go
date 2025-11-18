package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/domain"
	"github.com/pakdeetammasuk654-dot/numberniceic-api/internal/core/ports"
)

// satNumRepoPostgres เป็น Adapter สำหรับ PostgreSQL
type satNumRepoPostgres struct {
	DB *sql.DB
}

// NewSatNumRepoPostgres คือ Constructor ที่รับ Database Connection เข้ามา
func NewSatNumRepoPostgres(db *sql.DB) ports.SatNumRepositoryPort {
	return &satNumRepoPostgres{DB: db}
}

// ====================================================================
// Implementations of ports.SatNumRepositoryPort
// ====================================================================

// GetAllSatNums ดึงข้อมูลทั้งหมดจากตาราง sat_nums
func (r *satNumRepoPostgres) GetAllSatNums() ([]domain.SatNum, error) {
	query := "SELECT char_key, sat_value FROM public.sat_nums"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var satNums []domain.SatNum
	for rows.Next() {
		var sn domain.SatNum
		// สแกนข้อมูลจาก Row ลงใน Domain Entity โดยตรง
		if err := rows.Scan(&sn.CharKey, &sn.SatValue); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		satNums = append(satNums, sn)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return satNums, nil
}

// GetSatNumByCharKey ดึงข้อมูล SatNum จาก CharKey ที่กำหนด
func (r *satNumRepoPostgres) GetSatNumByCharKey(key string) (*domain.SatNum, error) {
	query := "SELECT char_key, sat_value FROM public.sat_nums WHERE char_key = $1"

	row := r.DB.QueryRow(query, key)

	sn := domain.SatNum{}
	// สแกนผลลัพธ์ลงใน struct
	if err := row.Scan(&sn.CharKey, &sn.SatValue); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Not Found
		}
		return nil, fmt.Errorf("query row scan failed: %w", err)
	}

	return &sn, nil
}
