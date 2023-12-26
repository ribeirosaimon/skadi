package repository

import (
	"gorm.io/gorm"
)

type pgsqlDatabase struct {
	databaseName string
	conn         *gorm.DB
}

var pgsqlDb *pgsqlDatabase

func (m pgsqlDatabase) Save(document Document) (Document, error) {
	
	return document, nil
}
