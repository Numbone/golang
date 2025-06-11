package db_sql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
}

func main() {

	db, err := sql.Open("postgres", "user=postgres password=admin dbname=postgres host=localhost port=5433 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	err = insertUser(db, User{
		Name:     "vasya",
		Email:    "petre@gmail.com",
		Password: "12312ad12",
	})

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	user, err := getUserById(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	fmt.Println(user)

}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Не удалось получить таблицы:", err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func getUserById(db *sql.DB, id int) (User, error) {
	var userId User
	err := db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&userId.ID, &userId.Name, &userId.Email, &userId.Password, &userId.RegisteredAt)
	return userId, err
}

func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO  logs(entity,action) VALUES ($1, $2)", "user", "created")
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteUserById(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func updateUser(db *sql.DB, u User, id int) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4", u.Name, u.Email, u.Password, id)
	if err != nil {
		return err
	}
	return nil
}
