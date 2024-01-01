package sqldomain

type User struct {
	Id   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (s User) GetId() uint64 {
	return s.Id
}

type UserTeste struct {
	Id   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (s UserTeste) GetId() uint64 {
	return s.Id
}
