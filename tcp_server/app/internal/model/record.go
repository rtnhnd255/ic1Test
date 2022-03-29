package model

type Record struct {
	PointID   int     `db:"_id"`
	DeviceID  string  `db:"device_id"`
	PointTime int     `db:"point_time"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Etc       string  `db:"etc"`
}

type RecordDTO struct {
	DeviceID  string  `db:"device_id"`
	PointTime int     `db:"point_time"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Etc       string  `db:"etc"`
}
