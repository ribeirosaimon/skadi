package sql

type Stock struct {
	Id   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (s Stock) GetId() uint64 {
	return s.Id
}
