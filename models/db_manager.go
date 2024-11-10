package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


// Custom Error Handling
//TODO: set up custom DBError and SQLError
// DBError for issues working with the database
// SQLError for errors running queries
// (might merge the two depending on what happens)
type SQLError struct {
  Code int16
  Err string
}

func (e *SQLError) Error() string {
  return fmt.Sprintf("SQLError Code %3d: %s", e.Code, e.Err)
} 

func NewSQLError(code int16, error string) error {
  return &SQLError{code,error}
}


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


// User CRUD Functions
func (db *DB) CreateUser(user *User) error {
  err := db.Get(user, "INSERT INTO users VALUES ($1, $2) RETURNING *", user.ID, user.Name)
  if err != nil {
    return NewSQLError(500, fmt.Sprintf("ERROR CREATING USER: %v", err))
  }
  return nil
}

// idfc anymore, i just want this to work now
// if user doesnt exist, ask user to register
func (db *DB) GetUserByEmail(email string) (*User, error) {
  var user User
  err := db.Get(&user, "SELECT * FROM users WHERE email = $1", email);
  if err == sql.ErrNoRows {
    return nil, NewSQLError(404, "USER NOT FOUND")
  } else if err != nil {
    return nil, NewSQLError(500, fmt.Sprintf("ERROR GETTING USER: %v", err))
  }
  return &user, nil
}

func (db *DB) UpdateUser(user *User) error {
  err := db.Get(user, "UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING *", user.Name, user.Email, user.ID)
  if err != nil {
    return NewSQLError(500, fmt.Sprintf("ERROR UPDATING USER (id=%v): %v", user.ID, err))
  }
  return nil
}

func (db *DB) DeleteUser(userID uuid.UUID) error {
  _, err := db.Exec("DELETE * FROM users WHERE id = $1", userID)
  if err != nil {
    return NewSQLError(500, fmt.Sprintf("ERROR DELETING USER (id=%v): %v", userID, err))
  }
  return nil
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
