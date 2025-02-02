package main

import (
	"interviewPrep/restAPI/db"
	"interviewPrep/restAPI/routes"
)

func main() {
	db.InitDB()
	r := routes.RouterConfig()
	r.Run(":8081")
}
