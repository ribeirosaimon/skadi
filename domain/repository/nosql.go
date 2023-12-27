package repository

import (
	"errors"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDatabase struct {
	databaseName string
	conn         *mongo.Database
}

var mongodb *mongoDatabase

func (m mongoDatabase) Save(document Document) (Document, error) {
	collection := getCollectionName(document)

	insertedID, err := m.conn.Collection(collection).InsertOne(nil, document)
	if err != nil {
		return nil, err
	}
	document.SetId(insertedID.InsertedID.(primitive.ObjectID))
	return document, nil
}

func (m mongoDatabase) FindById(document Document, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	collection := getCollectionName(document)
	if err := m.conn.Collection(collection).FindOne(nil, filter).Decode(document); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (m mongoDatabase) Find(document Document, filter bson.D) []Document {
	collection := getCollectionName(document)

	slice := reflect.SliceOf(reflect.TypeOf(document))
	sliceDocument := reflect.New(slice)

	cur, err := m.conn.Collection(collection).Find(nil, filter)
	if err != nil {
		panic(err)
	}
	err = cur.All(nil, sliceDocument.Interface())
	if err != nil {
		panic(err)
	}

	result := make([]Document, sliceDocument.Elem().Len())
	for i := 0; i < sliceDocument.Elem().Len(); i++ {
		result[i] = sliceDocument.Elem().Index(i).Interface().(Document)
	}

	return result
}

func (m mongoDatabase) DeleteById(id primitive.ObjectID, collection string) {
	filter := bson.D{{"_id", id}}
	m.conn.Collection(collection).DeleteOne(nil, filter)
}

func getCollectionName(document Document) string {
	documentType := reflect.TypeOf(document).Elem()
	name := documentType.Name()
	return strings.ToLower(name[:1]) + name[1:]
}
