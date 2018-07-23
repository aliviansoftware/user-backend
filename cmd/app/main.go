package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// The User Type (more like an object)
type User struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    County string `json:"County,omitempty"`
}

var users []User

// Display all from the users var
func GetUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

// Display a single data
func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range users {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&User{})
}

// create a new item
func CreateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var User User
    _ = json.NewDecoder(r.Body).Decode(&User)
    User.ID = params["id"]
    users = append(users, User)
    json.NewEncoder(w).Encode(users)
}

// Delete an item
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range users {
        if item.ID == params["id"] {
            users = append(users[:index], users[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(users)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    users = append(users, User{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", County: "County X"}})
    users = append(users, User{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", County: "County Y"}})
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
