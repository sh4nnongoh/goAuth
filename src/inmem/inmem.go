package inmem

import (
	"sync"
)

type userRepository struct {
	mtx sync.RWMutex
	users map[user.ID]*user.User
}

func (r *userRepository) Store(u *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[u.ID] = u
	return nil
}

func (r *userRepository) Find(id user.ID) (*user.User, error) {
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