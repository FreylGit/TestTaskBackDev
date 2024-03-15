package model

import (
	"time"
)

type RefreshToken struct {
	Id    string
	Token string
	Exp   time.Time
}
