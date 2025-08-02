package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rupak26/RESTapis_in_GOlang/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(ctg config.Config)(*Sqlite , error) {
	db , err := sql.Open("sqlite3" , ctg.StoragePath)
	if err != nil {
		return nil , err 
	}

	res , err := db.Exec(`CREATE TABLE IF NOT EXISTS students(
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   name TEXT,
	   email TEXT,
	   age INTEGER
	)`)
    fmt.Print(res)
	if err != nil {
		return nil , err
	} 

	return &Sqlite{
		Db : db ,
	} , nil 
}