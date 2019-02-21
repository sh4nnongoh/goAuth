package inmem

import (
	"sync"

	"github.com/sh4nnongoh/goAuth/src/user"
)

// NewUserRepository returns a new instance of a in-memory user repository.
func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[user.UserID]*user.User),
	}
}

type userRepository struct {
	mtx   sync.RWMutex
	users map[user.UserID]*user.User
}

func (r *userRepository) Store(u *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[u.UserID] = u
	return nil
}

func (r *userRepository) Find(id user.UserID) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.users[id]; ok {
		return val, nil
	}
	return nil, user.ErrUnknown
}

func (r *userRepository) FindAll() []*user.User {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	u := make([]*user.User, 0, len(r.users))
	for _, val := range r.users {
		u = append(u, val)
	}
	return u
}
