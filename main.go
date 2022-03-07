package main

import (
	"fmt"

	"github.com/armuh16/patternsingleton/repository"
	"github.com/armuh16/patternsingleton/singleton"
)

func main() {
	db := singleton.GetDBInstance()
	fmt.Println("from main", *singleton.GetDBInstance())
	result := db.GetUser("armuh")
	db.Status = "main"
	fmt.Println(db.Status, db.Type)
	fmt.Println(result)
	singleton.Umum = "hello"
	// singleton.instance = &db{}

	repository.HelloUser()
}
