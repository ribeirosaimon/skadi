package sql

type Session struct {
	Id     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId uint64 `json:"userId"`
	Token  string `json:"token"`
	BasicSQL
}

func (s Session) GetId() uint64 {
	return s.Id
}
