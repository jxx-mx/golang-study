package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    db, err := sql.Open("mysql","root:1234@tcp(localhost:3306)/sqlStudy")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() // flush 같은 작업, defer 를 사용함으로 어떠한 문제가 생겨도 db 커넥션은 종료

   err = db.Ping() // 핑을 날림으로써 db 가 정상 작동하는지 확인
   if err != nil {
        log.Fatal(err)
    }
    http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            getUsers(db, w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

		http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
					getUser(db, w, r)
				case "POST":
					createUser(db, w, r)
				case "PUT":
					updateUser(db, w, r)
				case "DELETE":
					deleteUser(db, w, r)

			default:
					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
	})



    fmt.Println("3000번 포트로 서버 실행")
    http.ListenAndServe(":3000", nil)
}

func getUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	fmt.Println(id)

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func getUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM user")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}


func createUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	result, err := db.Exec("INSERT INTO user (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	id, err := result.LastInsertId()
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	user.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	_, err = db.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	user.ID, err = strconv.Atoi(id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusNoContent)
}
