package repository

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/stretchr/testify/assert"
)

func TestSQLDatabase_CRUD(t *testing.T) {
	properties := config.GetProperties("test")
	repositoryTest := NewSkadiRepository(properties)

	t.Run("Save one value", func(t *testing.T) {
		entityStruct := newSimpleEntityStruct()
		err := repositoryTest.SqlTemplate().Save(&entityStruct)

		assert.Nil(t, err)
		assert.NotNil(t, entityStruct.GetId())
	})

	t.Run("Save one value with transactional", func(t *testing.T) {

		var idFound uint64

		err := repositoryTest.SqlTemplate().Transactional(func() error {
			// If I save struct
			entityStruct := newSimpleEntityStruct()
			err := repositoryTest.SqlTemplate().Save(&entityStruct)
			idFound = entityStruct.GetId()

			// I have to find it
			newEntityStructs := simpleEntityStruct{}
			findErr := repositoryTest.SqlTemplate().FindById(&newEntityStructs, entityStruct.GetId())

			// Both need to be equal
			assert.Nil(t, err)
			assert.Nil(t, findErr)
			assert.Equal(t, entityStruct.GetId(), newEntityStructs.GetId())

			// But if I return a error
			return errors.New("no transaction")
		})

		// This struct can not exist
		notFoundStruct := simpleEntityStruct{}
		findErr := repositoryTest.SqlTemplate().FindById(&notFoundStruct, idFound)

		assert.NotNil(t, err)
		assert.NotNil(t, findErr)
	})

	t.Run("Find by Id", func(t *testing.T) {
		entityStruct := newSimpleEntityStruct()
		repositoryTest.SqlTemplate().Save(&entityStruct)

		newEntityStructs := simpleEntityStruct{}

		err := repositoryTest.SqlTemplate().FindById(&newEntityStructs, entityStruct.GetId())

		assert.Nil(t, err)
		assert.Equal(t, entityStruct.GetId(), newEntityStructs.GetId())
		assert.Equal(t, entityStruct.Name, newEntityStructs.Name)
	})

	t.Run("Find by Id not found", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(1001)

		newEntityStructs := simpleEntityStruct{}
		err := repositoryTest.SqlTemplate().FindById(&newEntityStructs, uint64(randomNumber))

		assert.NotNil(t, err)
	})
}

type simpleEntityStruct struct {
	Id    uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name  string `json:"name"`
	Error bool   `json:"error" gorm:"not null"`
}

func newSimpleEntityStruct() simpleEntityStruct {
	return simpleEntityStruct{
		Name:  "Simple Data",
		Error: false,
	}
}

func (s *simpleEntityStruct) GetId() uint64 {
	return s.Id
}
