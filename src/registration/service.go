package registration

import (
	"github.com/sh4nnongoh/goAuth/src/user"
)

func NewService(userRepository user.Repository) Service {
	return service{UserRepository: userRepository}
}

type Service interface {
	New(user.UserID, user.Email, user.Name)
}

type service struct {
	UserRepository user.Repository
}

func (s service) New(id user.UserID, email user.Email, name user.Name) {
	user := user.User{UserID: id, Email: email, Name: name}
	s.UserRepository.Store(&user)
}
