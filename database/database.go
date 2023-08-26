package database

type IDatabase interface {
	Create(value interface{}) error
	CreateInBatches(value interface{}, batchSize int) error
	Update(value interface{}) error
	FindOne(dest interface{}, opts ...FindOneOptions) error
	FindMany(dest interface{}, opts ...FindManyOptions) error
	Delete(value interface{}, opts ...DeleteOptions) error
	Count(count *int64) error
}

type Query struct {
	query string
	args  []interface{}
}

type FindManyOptions struct {
	offset int64
	limit  int64
	sort   interface{}
	query  []Query
}

type FindOneOptions struct {
	sort  interface{}
	query []Query
}

type DeleteOptions struct {
	query []Query
}
