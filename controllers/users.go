package controllers

import (
	"net/http"

	"github.com/didiyudha/go-mvc-pattern/models"
	"github.com/julienschmidt/httprouter"
)

// UserController Struct
type UserController struct{}

var usr = models.NewUser()

// NewUserController return pointer UserController
func NewUserController() *UserController {
	return &UserController{}
}

// GetAllUsers grab all users from database
func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := usr.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "users.gohtml", users)
}

// CreateUser create a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	u.Username = r.FormValue("uname")
	u.FirstName = r.FormValue("fname")
	u.LastName = r.FormValue("lname")
	bPass, err := u.EncryptPassword(r.FormValue("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Password = bPass
	err = usr.Save(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
