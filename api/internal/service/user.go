package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ribeirosaimon/skadi/api/internal/repository"
	"github.com/ribeirosaimon/skadi/api/internal/util"
	"github.com/ribeirosaimon/skadi/domain/sql"
)

var userService = &UserService{}

type UserService struct{}

func (s *UserService) SignUp(name, familyName, email, password string) (sql.User, error) {
	var user sql.User
	if name == "" {
		return user, errors.New("name can't be null")
	}
	if email == "" {
		return user, errors.New("email can't be null")
	}
	if password == "" {
		return user, errors.New("password can't be null")
	}
	encryptPassword, err := util.EncryptPassword(password)
	if err != nil {
		return user, err
	}
	now := time.Now()

	user.Name = name
	user.Email = email
	user.FamilyName = familyName
	user.Password = encryptPassword
	user.Status = sql.UserActive
	user.CreatedAt = now
	user.UpdatedAt = now
	if err := repository.GetRepository().SqlTemplate().Save(user); err != nil {
		return sql.User{}, err
	}
	return user, nil
}

func (s *UserService) SignIn(email, password string) (sql.User, error) {
	var user sql.User
	if err := repository.GetRepository().SqlTemplate().
		CreateNativeQuery(fmt.Sprintf("select u.* from user where u.email = %s", email), &user); err != nil {
		return sql.User{}, err
	}
	return sql.User{}, nil
}

func NewUserService() *UserService {
	return userService
}
