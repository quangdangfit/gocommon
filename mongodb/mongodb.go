package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	SortDescending = -1
	_idField       = "_id"
)

type DB interface {
	FindById(ctx context.Context, collection string, id string, result interface{}) error
	FindOne(ctx context.Context, collection string, result interface{}, opts ...Option) error
	Find(ctx context.Context, collection string, result interface{}, opts ...Option) error
	UpdateMany(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error
	UpdateOne(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error
	UpdateOneRaw(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error
	UpdateManyRaw(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error
	Insert(ctx context.Context, collection string, document interface{}) error
	GetCollection(ctx context.Context, collection string) *mongo.Collection
	EnsureIndexes(ctx context.Context, collection string, model []mongo.IndexModel) error
	Delete(ctx context.Context, collection string, filter interface{}, opts *options.DeleteOptions) error
	WithTransaction(callback func(sc mongo.SessionContext) (interface{}, error)) (interface{}, error)
	FindOneAndUpdate(ctx context.Context, collection string, filter, update interface{}, opts *options.FindOneAndUpdateOptions, result interface{}) error
	Count(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error)
	BulkWriteRaw(ctx context.Context, collection string, operations []mongo.WriteModel, opts ...*options.BulkWriteOptions) error
	Close()
}

type updateFn func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

func New(config *Config) (DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	if config.Database == "" {
		return nil, errors.New("no db name detected")
	}

	db := client.Database(config.Database)

	return &dbImpl{
		db: db,
	}, nil
}
