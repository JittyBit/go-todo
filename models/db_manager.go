package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Custom Error Handling
//TODO: set up custom DBError and SQLError
// DBError for issues working with the database
// SQLError for errors running queries
// (might merge the two depending on what happens)

// Set up database connection layer
type DB struct {
	*sqlx.DB
}


//TODO: Maybe return a DBError here instead
func NewDB(connectionString string) (*DB, error) {
	db, err := sqlx.Open("sqlite3", connectionString)
	if err != nil {
		return nil, fmt.Errorf("DBError: ERROR OPENING NEW DATABASE: %w", err)
	}

  if err = db.Ping(); err != nil {
    return nil, fmt.Errorf("DBError: ERROR CONNECTING TO DATABASE: %w", err)
  }

	return &DB{db}, nil
}

/*
database/sql steps:
step   1: query -> prepared statement
step 1.5: defer close statement
step   2: run pstmt.Query() with args
step   3: iterate over rows
step 3.5: defer close rows
step   4: check rows.Err()

REMEMBER: use Exec() instead if query doesnt return rows
- INSERT, UPDATE, DELETE

single row query:
var (
  id int
  name string
)
err := db.QueryRow("select name from users where id = ?", 1).Scan(&name)
pstmt, _ := db.Prepare("<insert query>")
err := pstmt.QueryRow(<insert query>).Scan(&id)

maybe use Transactions, but idk when yet
possibly for QueryTable-like reasons
*/
