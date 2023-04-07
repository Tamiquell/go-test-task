package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github/Tamiquell/go-test-task/api"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println(("adding admins and one user to BD"))
	admin_id_1 := uuid.New().String()
	admin_id_2 := uuid.New().String()
	user_id_1 := uuid.New().String()
	api.UsersDB[admin_id_1] = api.User{
		ID:       admin_id_1,
		Email:    "",
		Username: "admin1",
		Password: "admin1",
		Admin:    true,
	}
	api.UsersDB[admin_id_2] = api.User{
		ID:       admin_id_2,
		Email:    "",
		Username: "admin2",
		Password: "admin2",
		Admin:    true,
	}
	api.UsersDB[user_id_1] = api.User{
		ID:       user_id_1,
		Email:    "",
		Username: "user",
		Password: "user",
		Admin:    false,
	}
}

func main() {

	r := mux.NewRouter()

	r = r.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/user", api.CreateUser).Methods("POST")
	r.HandleFunc("/user", api.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", api.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	fmt.Println("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server stopped")
}
