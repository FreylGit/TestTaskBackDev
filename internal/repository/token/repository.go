package token

import (
	"context"
	"github.com/FreylGit/TestTaskBackDev/internal/client/db"
	"github.com/FreylGit/TestTaskBackDev/internal/converter"
	"github.com/FreylGit/TestTaskBackDev/internal/model"
	"github.com/FreylGit/TestTaskBackDev/internal/repository"
	converterRepo "github.com/FreylGit/TestTaskBackDev/internal/repository/token/converter"
	modelRepo "github.com/FreylGit/TestTaskBackDev/internal/repository/token/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "Auth"
	collectionName = "Token"
	idColumn       = "_id"
	tokenColumn    = "token"
	expColumn      = "exp"
)

type repo struct {
	client *db.Client
	c      *mongo.Collection
}

func NewRepository(client *db.Client) repository.TokenRepository {
	c := client.DB.Database(databaseName).Collection(collectionName)

	return &repo{client: client, c: c}
}

func (r *repo) Create(ctx context.Context, rtoken *model.RefreshToken) error {
	rtokenRepo := converter.ToTokenCreateFromService(*rtoken)
	_, err := r.c.InsertOne(ctx, rtokenRepo)
	if err != nil {
		return repository.ErrCreate
	}

	return nil
}

func (r *repo) Get(ctx context.Context, token string) (*model.RefreshToken, error) {
	filter := bson.D{{tokenColumn, token}}
	var tokenModel modelRepo.RefreshToken
	err := r.c.FindOne(ctx, filter).Decode(&tokenModel)
	if err != nil {
		return nil, repository.ErrNotFound
	}

	return converterRepo.ToTokenFromRepo(tokenModel), nil
}

func (r *repo) Delete(ctx context.Context, token string) error {
	filter := bson.D{{tokenColumn, token}}
	result, err := r.c.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *repo) Update(ctx context.Context, token *model.RefreshToken) error {
	id, err := primitive.ObjectIDFromHex(token.Id)
	filter := bson.D{{idColumn, id}}
	update := bson.D{{"$set", bson.D{{tokenColumn, token.Token}, {expColumn, token.Exp}}}}

	result, err := r.c.UpdateOne(ctx, filter, update)
	if err != nil {
		return repository.ErrUpdate
	}
	if result.ModifiedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}
