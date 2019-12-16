/*
File describe point structure.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
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
