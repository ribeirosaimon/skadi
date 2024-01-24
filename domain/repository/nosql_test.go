package repository

import (
	"testing"

	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoDatabase_Utils(t *testing.T) {
	name := getCollectionName(&simpleDataStruct{})
	assert.Equal(t, "simpleDataStruct", name)
}

func TestMongoDatabase_CRUD(t *testing.T) {
	properties := config.GetPropertiesFile("test")

	repositoryTest := NewSkadiRepository(properties)

	t.Run("Save one value", func(t *testing.T) {
		// Given
		dataStruct := newSimpleDataStruct()
		// When
		savedValue, err := repositoryTest.NoSqlTemplate().Save(&dataStruct)

		// Verification
		var newDataStruct simpleDataStruct
		repositoryTest.NoSqlTemplate().FindById(&newDataStruct, savedValue.GetId())

		assert.Nil(t, err)
		assert.NotEmpty(t, savedValue.GetId())
		assert.Equal(t, savedValue.GetId(), newDataStruct.GetId())
	})

	t.Run("Find by Id", func(t *testing.T) {
		// Given
		dataStruct := newSimpleDataStruct()
		// When
		savedValue, _ := repositoryTest.NoSqlTemplate().Save(&dataStruct)

		var newDataStruct simpleDataStruct
		err := repositoryTest.NoSqlTemplate().FindById(&newDataStruct, savedValue.GetId())

		// Verification
		assert.Nil(t, err)
		assert.NotEmpty(t, savedValue.GetId())
		assert.Equal(t, savedValue.GetId(), newDataStruct.GetId())
	})

	t.Run("Not found when pass wrong Id", func(t *testing.T) {
		var newDataStruct simpleDataStruct
		err := repositoryTest.NoSqlTemplate().FindById(&newDataStruct, primitive.NewObjectID())

		// Verification
		assert.NotNil(t, err)
		assert.Equal(t, "mongo: no documents in result", err.Error())
	})

	t.Run("Find All", func(t *testing.T) {
		// Given
		newName := primitive.NewObjectID()

		dataStruct := newSimpleDataStruct()
		dataStruct2 := newSimpleDataStruct()

		dataStruct.Name = newName.Hex()
		dataStruct2.Name = newName.Hex()

		repositoryTest.NoSqlTemplate().Save(&dataStruct)
		repositoryTest.NoSqlTemplate().Save(&dataStruct2)

		// When
		find := repositoryTest.NoSqlTemplate().Find(&simpleDataStruct{}, bson.D{{"name", newName.Hex()}})

		// Verification
		assert.Equal(t, len(find), 2)
	})

	t.Run("Delete one value", func(t *testing.T) {
		// Given
		dataStruct := newSimpleDataStruct()
		// When
		savedValue, _ := repositoryTest.NoSqlTemplate().Save(&dataStruct)

		repositoryTest.NoSqlTemplate().DeleteById(savedValue.GetId(), "simpleDataStruct")

		var newDataStruct simpleDataStruct
		notFoundError := repositoryTest.NoSqlTemplate().FindById(&newDataStruct, savedValue.GetId())

		assert.NotNil(t, notFoundError)
		assert.Equal(t, "mongo: no documents in result", notFoundError.Error())
	})

}

type simpleDataStruct struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func newSimpleDataStruct() simpleDataStruct {
	return simpleDataStruct{Name: "Simple Data"}
}

func (s *simpleDataStruct) GetId() primitive.ObjectID {
	return s.Id
}

func (s *simpleDataStruct) SetId(id primitive.ObjectID) {
	s.Id = id
}
