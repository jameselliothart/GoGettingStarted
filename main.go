package main

import (
	"fmt"

	"github.com/jameselliothart/GoGettingStarted/models"
)

func main() {
	u := models.User{
		Id:        2,
		FirstName: "Tricia",
		LastName:  "McMillian",
	}
	fmt.Println(u)
}
