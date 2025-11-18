package domain

// SatNum คือ Entity ที่แทนข้อมูลจากตาราง public.sat_nums
// ซึ่งเก็บคู่ CharKey และ SatValue
type SatNum struct {
	// CharKey: คอลัมน์ char_key (text) ในฐานข้อมูล
	CharKey string `json:"char_key"`

	// SatValue: คอลัมน์ sat_value (integer) ในฐานข้อมูล
	SatValue int `json:"sat_value"`
}
