package auth

import (
	"crypto/rand"
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func NewAuthService() AuthService {
	return authService{}
}

type AuthService interface {
	Authorize()
	AccessToken()
	ValidateUser(username string, password string)
}

type authService struct {
	Authcode Authcode
}

func (s authService) Authorize() {
	s.GenerateAuthCode()
}

func (s authService) AccessToken() {

}

func (s authService) ValidateUser() {

}

func (s *authService) GenerateAuthCode() {
	b := make([]byte, 50)
	rand.Read(b)

	s.Authcode.ID = bson.NewObjectId()
	s.Authcode.Code = uuid.Must(uuid.NewV4()).String()
	//s.Authcode.UserID = user.ID
	s.Authcode.ExpiresAt = time.Now().Local().Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(1) +
		time.Second*time.Duration(0))
	s.Authcode.CreatedAt = time.Now().Local()
	s.Authcode.UpdatedAt = time.Now().Local()

}

type Authcode struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Code      string        `bson:"code" json:"code"`
	ExpiresAt time.Time     `bson:"expires_at" json:"expires_at"`
	UserID    bson.ObjectId `bson:"user_id" json:"user_id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
