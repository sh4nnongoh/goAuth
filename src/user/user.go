package user

// New creates a new user
func New(userID UserID, email Email, name Name) *User {
	return &User{UserID: userID, Email: email, Name: name}
}

// UserID uniquely identifies a user.
type UserID string
type Email string
type Name string

// User is the central class in the domain model.
type User struct {
	UserID UserID
	Email  Email
	Name   Name
}

type Repository interface {
	Store(user *User) error
	Find(id UserID) (*User, error)
	FindAll() []*User
}
