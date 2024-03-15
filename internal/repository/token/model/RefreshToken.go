package model

import (
	"time"
)

type RefreshToken struct {
	Id    string    `bson:"_id"`
	Token string    `bson:"token"`
	Exp   time.Time `bson:"exp"`
}

type RefreshTokenCreate struct {
	Token string    `bson:"token"`
	Exp   time.Time `bson:"exp"`
}
