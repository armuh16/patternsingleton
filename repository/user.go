package repository

import (
	"fmt"

	"github.com/armuh16/patternsingleton/singleton"
)

func HelloUser() {
	db := singleton.GetDBInstance()
	fmt.Println("from user", *singleton.GetDBInstance())
	db.Status = "user"
	fmt.Println(db.GetUser("armuh"), db.Status, db.Type)
}
