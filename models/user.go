package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("user already exists")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(ID int) (User, error) {
	for _, user := range users {
		if user.ID == ID {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", ID)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v is not found", u.ID)
}

func RemoveUserByID(ID int) error {
	for i, u := range users {
		if u.ID == ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v is not found", ID)
}
