package main

import (
	_ "car/ent/runtime"
	"car/gateways/sql_db"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	db := sql_db.SqlDb{}
	db.Migrate()
	defer db.Close()
}
