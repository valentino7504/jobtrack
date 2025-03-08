package main

import (
	"fmt"

	"github.com/valentino7504/jobtrack/cmd"
	"github.com/valentino7504/jobtrack/internal/db"
)

func main() {
	sqliteDB, err := db.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.InitDB(sqliteDB); err != nil {
		fmt.Println(err)
		return
	}
	cmd.SetDB(sqliteDB)
	cmd.Execute()
}
