package controllers

import "net/http"

// RegisterControllers : Takes care of the front urls
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
