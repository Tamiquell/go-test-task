package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if !UserIsAuthozied(r) {
		http.Error(w, "You must be authorized to use this command", http.StatusForbidden)
		return
	}
	if !UserIsAdmin(r) {
		http.Error(w, "You must be admin to use this command", http.StatusForbidden)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = uuid.New().String()
	UsersDB[user.ID] = user
	w.WriteHeader(http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	if !UserIsAuthozied(r) {
		http.Error(w, "You must be authorized to use this command", http.StatusForbidden)
		return
	}

	var usersList []User
	for _, user := range UsersDB {
		usersList = append(usersList, user)
	}
	json.NewEncoder(w).Encode(usersList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !UserIsAuthozied(r) {
		http.Error(w, "You must be authorized to use this command", http.StatusForbidden)
		return
	}

	id := mux.Vars(r)["id"]
	user, ok := UsersDB[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if !UserIsAuthozied(r) {
		http.Error(w, "You must be authorized to use this command", http.StatusForbidden)
		return
	}
	if !UserIsAdmin(r) {
		http.Error(w, "You must be admin to use this command", http.StatusForbidden)
		return
	}

	id := mux.Vars(r)["id"]
	user, ok := UsersDB[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UsersDB[id] = user
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !UserIsAuthozied(r) {
		http.Error(w, "You must be authorized to use this command", http.StatusForbidden)
		return
	}
	if !UserIsAdmin(r) {
		http.Error(w, "You must be admin to use this command", http.StatusForbidden)
		return
	}

	id := mux.Vars(r)["id"]
	_, ok := UsersDB[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(UsersDB, id)
	w.WriteHeader(http.StatusNoContent)
}

func UserIsAdmin(r *http.Request) bool {
	request_user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}
	for _, u := range UsersDB {
		if u.Username == request_user && u.Password == pass && u.Admin {
			return true
		}
	}
	return false
}

func UserIsAuthozied(r *http.Request) bool {
	request_user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}
	for _, u := range UsersDB {
		if u.Username == request_user && u.Password == pass {
			return true
		}
	}
	return false
}
