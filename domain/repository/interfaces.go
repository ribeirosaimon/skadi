package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entity is sql domains
type Entity interface {
	GetId() uint64
}

// Document is no-Sql domains
type Document interface {
	GetId() primitive.ObjectID
	SetId(id primitive.ObjectID)
}

// skadiRepositoryInterface is an interface between no-sql & sql
type skadiRepositoryInterface interface {
	NoSqlTemplate() noSqlTemplateInterface
	SqlTemplate() sqlTemplateInterface
}

// sqlTemplateInterface is a sql database interface
type sqlTemplateInterface interface {
	Save(Entity) error
	FindById(Entity, uint64) error
	FindAll() []Entity
	Transactional(func() error) error
}

// noSqlTemplateInterface is a no-sql database interface
type noSqlTemplateInterface interface {
	Save(Document) (Document, error)
	FindById(Document, primitive.ObjectID) error
	Find(Document, bson.D) []Document
	DeleteById(primitive.ObjectID, string)
}
