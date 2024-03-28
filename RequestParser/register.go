package requestparser

import (
	"AnimeList/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Email    string
	Username string
	Password string
}

func (r Register) ParseRegister() (model.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(r.Password), 14)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:       uuid.New(),
		Email:    r.Email,
		Username: r.Username,
		Password: string(hashedPass),
	}

	return user, nil
}
