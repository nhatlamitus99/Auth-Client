package entities

import (
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      string `json:"age"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPassword: %s\n", user.Id, user.Name, user.Password)
}
