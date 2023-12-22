package repository

import (
	"fmt"

	"github.com/magiconair/properties"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type skadiRepository struct {
	conn *gorm.DB
}

func NewSkadiRepository(p *properties.Properties) skadiInterface {
	repository := &skadiRepository{}
	repository.connectSqlDatabase(p)
	return repository
}
func (s *skadiRepository) FindById(any) (any, error) {
	return "", nil
}

func (s *skadiRepository) connectSqlDatabase(p *properties.Properties) {
	dbUsername := p.GetString("database.username", "")
	dbPassword := p.GetString("database.password", "")
	dbName := p.GetString("database.name", "")
	dbPort := p.GetInt("database.port", 0)
	dbHost := p.GetString("database.host", "")
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%d sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	if s.conn == nil {
		var err error
		s.conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
}
