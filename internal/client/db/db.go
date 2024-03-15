package db

import "go.mongodb.org/mongo-driver/mongo"

type DB interface {
}

type Client struct {
	DB *mongo.Client
}

func NewClient(db *mongo.Client) *Client {
	return &Client{DB: db}
}
