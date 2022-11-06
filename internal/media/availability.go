package media

import "time"

type Availability struct {
	LastOnlineAt time.Time `gorm:"column:last_online_at"`
	PeerCount    int       `gorm:"column:peer_count"`
}
