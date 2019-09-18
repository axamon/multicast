package multicast

import "time"

type Archivio struct {
	Index      int       `json:"index"`
	Aggiornato bool      `json:"aggiornato"`
	Timestamp  time.Time `json:"timestamp"`
}
