package dto

import "golang.org/x/crypto/bcrypt"

type Student struct {
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Password         string   `json:"password"`
	Description      string   `json:"description,omitempty"`
	University       string   `json:"university,omitempty"`
	Degree           string   `json:"degree,omitempty"`
	Skills           []string `json:"skills,omitempty"`
	YearOfGraduation int      `json:"year_of_graduation,omitempty"`
}

func (u *Student) HashAndSalt() *Student {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hash)
	return u
}

func NewStudent() *Student {
	return &Student{}
}
