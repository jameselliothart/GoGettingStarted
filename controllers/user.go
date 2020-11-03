package controllers

import (
	"regexp"
	"net/http"
)

type userController struct{
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi from user controller"))
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
