package repository

import (
	"gorm.io/gorm"
)

type pgsqlDatabase struct {
	databaseName string
	conn         *gorm.DB
}

func (m *pgsqlDatabase) CreateNativeQuery(query string, response interface{}) error {
	if tx := m.conn.Raw(query).Scan(response); tx.Error != nil {
		return tx.Error
	}
	return nil

}

var pgsqlDb *pgsqlDatabase

func (m *pgsqlDatabase) FindAll(entity interface{}) error {
	if err := m.conn.Find(entity).Error; err != nil {
		return err
	}

	return nil
}

func (m *pgsqlDatabase) FindById(entity Entity, id uint64) error {
	if err := m.conn.First(entity, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *pgsqlDatabase) Save(entity Entity) error {
	m.entityMigrate(entity)

	if err := m.conn.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (m *pgsqlDatabase) Transactional(myFunc func() error) error {
	m.conn.Begin()
	m.conn.SavePoint("begin")

	err := myFunc()

	if err != nil {
		return m.conn.RollbackTo("begin").Error
	}
	return m.conn.Commit().Error
}

// Deprecated
func (m *pgsqlDatabase) entityMigrate(entity Entity) {
	if err := m.conn.AutoMigrate(entity); err != nil {
		panic("error")
	}
}
