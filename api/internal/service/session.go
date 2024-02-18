package service

import (
	"time"

	"github.com/ribeirosaimon/skadi/api/internal/repository"
	"github.com/ribeirosaimon/skadi/api/internal/util"
	"github.com/ribeirosaimon/skadi/domain/sql"
)

var sessionService = &SessionService{}

type SessionService struct{}

func (s *SessionService) CreateSession(user sql.User) {
	var session sql.Session

	session.Token = util.GenerateUUIDToken()
	session.UserId = user.Id
	now := time.Now()
	session.CreatedAt = now
	session.UpdatedAt = now

	repository.GetRepository().SqlTemplate().Save(session)
}
