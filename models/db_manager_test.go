package models_test

import (
	"log"
	"testing"

	DBMan "github.com/JittyBit/go-todo/models"
)

func TestOpen(t *testing.T) {
	db, err := DBMan.NewDB("path/to/sqlite3/db")
  if err != nil {
    log.Fatalln(err.Error())
  }
  defer db.Close()
}
