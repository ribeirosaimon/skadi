package repository

import (
	"os"

	"github.com/ribeirosaimon/skadi/api/internal/skadiEngine"
	"github.com/ribeirosaimon/skadi/domain/repository"
)

var repo repository.SkadiRepositoryInterface

func GetRepository() repository.SkadiRepositoryInterface {
	return repo
}

func init() {
	env := os.Args[1]
	repo = repository.NewSkadiRepository(skadiEngine.GetProperties(env))
}
