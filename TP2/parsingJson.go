package main

import (
	"fmt"
	"os"
)

type User struct {
	Login string
	Password string
}

func newUser(Login string) *user {
	u.Login := User{Login: Login}
	u.Password :=  User{Password: Password}
	return &u
}

func main () {
	fmt.Println("Le nom d'utilisateur est ",username)
	fmt.Println("Le mot de passe est ",password)
}
