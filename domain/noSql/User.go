package noSql

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

func (u *User) GetId() primitive.ObjectID {
	return u.Id
}

func (u *User) SetId(id primitive.ObjectID) {
	u.Id = id
}
