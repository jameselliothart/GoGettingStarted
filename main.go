package main

import (
	"net/http"

	"github.com/jameselliothart/GoGettingStarted/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
