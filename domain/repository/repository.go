package repository

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/skadi/domain/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SkadiRepository struct {
	sqlConn   sqlTemplateInterface
	noSqlConn noSqlTemplateInterface
}

func NewSkadiRepository(p *properties.Properties) SkadiRepositoryInterface {
	repository := &SkadiRepository{}

	repository.connectNoSqlDatabase(p)
	repository.connectSqlDatabase(p)

	return repository
}

func (s *SkadiRepository) NoSqlTemplate() noSqlTemplateInterface {
	return mongodb
}

func (s *SkadiRepository) SqlTemplate() sqlTemplateInterface {
	return pgsqlDb
}

func (s *SkadiRepository) connectNoSqlDatabase(p *properties.Properties) {

	mongoUrl := p.GetString("database.mongo.url", "")
	dbName := p.GetString("database.mongo.name", "")

	mongodb = &mongoDatabase{databaseName: dbName}

	connection, _ := mongo.Connect(nil, options.Client().ApplyURI(mongoUrl+dbName))

	if err := connection.Ping(nil, nil); err != nil {
		panic(err)
	}
	mongodb.conn = connection.Database(mongodb.databaseName)
}

func (s *SkadiRepository) connectSqlDatabase(p *properties.Properties) {
	dbUsername := p.GetString("database.username", "")
	dbPassword := p.GetString("database.password", "")
	dbName := p.GetString("database.name", "")
	dbPort := p.GetInt("database.port", 0)
	dbHost := p.GetString("database.host", "")
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%d sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	pgsqlDb = &pgsqlDatabase{}
	var err error
	if strings.Contains(dsn, "test") {
		var file string
		inMemoryDb := p.GetBool("database.test.in-memory", false)
		if !inMemoryDb {
			dir, _ := config.FindRootDir()
			file = filepath.Join(dir, p.GetString("database.host", ""))
		} else {
			file = p.GetString("database.in-memory.host", "")
		}
		pgsqlDb.conn, err = gorm.Open(sqlite.Open(file), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic("erro connection Db")
		}
	} else {
		pgsqlDb.conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
	}
	if err != nil {
		panic(err)
	}
	// s.migrateValues(pgsqlDb.conn)
}
