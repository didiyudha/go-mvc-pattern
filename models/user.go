package models

import (
	"log"

	"github.com/didiyudha/go-mvc-pattern/db"
	"golang.org/x/crypto/bcrypt"
)

const (
	ins = `INSERT INTO users (username, first_name, last_name, password)
				 VALUES(?, ?, ?, ?)`
)

// User struct
type User struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	Password  []byte
}

// NewUser return pointer to User and initialize it
func NewUser() *User {
	return &User{}
}

// FindAll get all users from database
func (u *User) FindAll() ([]User, error) {
	users := make([]User, 0)
	dbs, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		return users, err
	}
	defer dbs.Close()
	rows, err := dbs.Query(`SELECT id, username, first_name, last_name FROM users`)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		u := User{}
		rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName)
		users = append(users, u)
	}
	return users, nil
}

//Save save a user to database
func (u *User) Save(newUsr User) error {
	dbs, err := db.GetDB()
	if err != nil {
		return err
	}
	defer dbs.Close()
	_, err = dbs.Exec(ins, newUsr.Username, newUsr.FirstName, newUsr.LastName, newUsr.Password)
	if err != nil {
		return err
	}
	return nil
}

// EncryptPassword generate new password from a registered user
func (u *User) EncryptPassword(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
}
