package models

import "time"

type Point struct {
	Client       uint32
	NavigateDate time.Time
	PacketID     float64
	Latitude     float64
	Longitude    float64
	Course       uint8
}
