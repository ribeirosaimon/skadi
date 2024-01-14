package repository

import (
	"github.com/ribeirosaimon/skadi/api/internal/skadiEngine"
	"github.com/ribeirosaimon/skadi/domain/repository"
)

var repo repository.SkadiRepositoryInterface

func GetRepository() repository.SkadiRepositoryInterface {
	return repo
}

func init() {
	repo = repository.NewSkadiRepository(skadiEngine.GetProperties())
}
