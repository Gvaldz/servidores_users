package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	application "usuarios/src/users/application"
	infrastructure "usuarios/src/users/infrastructure"
	entities "usuarios/src/users/domain/entities"
)

var userRepo = infrastructure.NewUserRepository()
var createUser = application.NewCreateUser(userRepo)
var updateUser = application.NewUpdateUser(userRepo)
var listUsers = application.NewListUsers(userRepo)
var deleteUser = application.NewDeleteUser(userRepo)

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUser)
	http.ListenAndServe(":8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := listUsers.Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user entities.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := createUser.Execute(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var user entities.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.ID = id
		if err := updateUser.Execute(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		if err := deleteUser.Execute(int32(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}