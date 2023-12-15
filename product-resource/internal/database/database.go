package database

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/rsfreitas/go-pocket-utils/errors"
	utlogger "github.com/rsfreitas/go-pocket-utils/logger"
)

type NewOptions struct {
	Port         int
	Username     string
	Password     string
	Hostname     string
	DatabaseName string
	Errors       *errors.Factory
}

type Database struct {
	databaseName string
	errors       *errors.Factory
	*gorm.DB
}

func New(options NewOptions) (*Database, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		options.Username, options.Password, options.Hostname,
		options.Port, options.DatabaseName)

	dbConn, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &Database{
		DB:           dbConn,
		errors:       options.Errors,
		databaseName: options.DatabaseName,
	}, nil
}

func (d *Database) FindOneByID(
	ctx context.Context,
	table, id string,
	target interface{},
) error {
	return d.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(table).First(target).Where("id == ?", id).Error; err != nil {
			return d.errors.NotFound().WithAttributes(
				utlogger.Error(err),
				utlogger.Any("database", d.databaseName),
				utlogger.Any("table", table),
			).Submit(ctx)
		}

		return nil
	})
}

func (d *Database) Insert(
	ctx context.Context,
	table string,
	source interface{},
) error {
	return d.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(table).Create(source).Error; err != nil {
			return d.errors.Internal(err).WithAttributes(
				utlogger.Any("database", d.databaseName),
				utlogger.Any("table", table),
			).Submit(ctx)
		}

		return nil
	})
}

func (d *Database) Update(
	ctx context.Context,
	table string,
	source interface{},
	target interface{},
) error {
	return d.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(table).Model(source).Updates(target).Error; err != nil {
			return d.errors.Internal(err).WithAttributes(
				utlogger.Any("database", d.databaseName),
				utlogger.Any("table", table),
			).Submit(ctx)
		}

		return nil
	})
}
