package sql

type UserStatus string

const (
	UserActive    UserStatus = "ACTIVE"
	UserInactive  UserStatus = "INACTIVE"
	UserSuspended UserStatus = "SUSPENDED"
)

type User struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string     `json:"name"`
	FamilyName string     `json:"familyName"`
	Email      string     `json:"email"`
	Password   string     `json:"-"`
	Roles      []Role     `gorm:"type:text[]" json:"roles"`
	Status     UserStatus `json:"status"`
	BasicSQL
}

func (s User) GetId() uint64 {
	return s.Id
}
