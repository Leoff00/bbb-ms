package voto

import (
	"time"
)

type Voto struct {
	Id        int       `json:"-"`
	Voto      int       `json:"voto"`
	CreatedAt time.Time `json:"createdAt"`
}
