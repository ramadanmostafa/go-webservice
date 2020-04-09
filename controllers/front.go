package controllers

import "net/http"

func RegisterControllers() {
	uc := CreateUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
