package http_server

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{
		Id:   "1",
		Name: "Jonh",
	},
	{
		Id:   "2",
		Name: "Atom",
	},
}

func server() {
	http.HandleFunc("/users", handleUser)
	if err := http.ListenAndServe(":5555", nil); err != nil {
		log.Fatal(err)
	}
}
func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}
	w.Write(resp)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 Created"))
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}
