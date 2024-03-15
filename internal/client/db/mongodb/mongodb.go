package mongodb

import (
	"github.com/FreylGit/TestTaskBackDev/internal/client/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	db *mongo.Client
}

func NewDb(db *mongo.Client) db.DB {
	return mongodb{db: db}
}
