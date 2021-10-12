package main

import (
	_ "database/sql"

	"github.com/milencium/reglogin/db"
	"github.com/milencium/reglogin/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	
	r.Run(":8081")
}
