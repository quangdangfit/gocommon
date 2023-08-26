package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/quangdangfit/gocommon/logger"
)

func NewDatabase(databaseURI string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return database
}

type postgresDB struct {
	db *gorm.DB
}

func (d *postgresDB) Create(value interface{}) error {
	return d.db.Create(value).Error
}

func (d *postgresDB) CreateInBatches(value interface{}, batchSize int) error {
	return d.db.CreateInBatches(value, batchSize).Error
}

func (d *postgresDB) Update(value interface{}) error {
	return d.db.Save(value).Error
}

func (d *postgresDB) FindOne(dest interface{}, opts ...FindOneOptions) error {
	return d.db.First(dest, opts).Error
}
