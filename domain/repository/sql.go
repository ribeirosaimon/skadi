package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type pgsqlDatabase struct {
	databaseName string
	conn         *gorm.DB
}

var pgsqlDb *pgsqlDatabase

func (m pgsqlDatabase) FindAll() []Entity {
	// TODO implement me
	panic("implement me")
}

func (m pgsqlDatabase) FindById(entity Entity, id uint64) error {
	if err := m.conn.First(entity, id).Error; err != nil {
		return err
	}
	return nil
}

func (m pgsqlDatabase) Save(entity Entity) error {
	m.entityMigrate(entity)

	if err := m.conn.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (m pgsqlDatabase) Transactional(myFunc func() error) error {
	if err := m.conn.Transaction(func(tx *gorm.DB) error {
		tx.Begin()

		if err := myFunc(); err != nil {
			rollback := tx.Rollback()
			fmt.Sprintf("%s", rollback.Error)
			return err
		}
		return tx.Commit().Error

	}); err != nil {
		return err
	}
	return nil
}

// Deprecated
func (m pgsqlDatabase) entityMigrate(entity Entity) {
	if err := m.conn.AutoMigrate(entity); err != nil {
		panic("error")
	}
}
