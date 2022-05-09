package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dashboard( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website, you've requested: %s\n", r.URL.Path)
}

func bookController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func CreateUserTable(db) {
	query := `
		CREATE TABLE IF NOT EXISTS users1 (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
	}
}

func main() {

	fmt.Println("--- start ---")

	db, err := sql.Open("mysql", "root:Abc@123654@(127.0.0.1:3308)/HelloSql?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected to MySQL database")
	if err := db.Ping(); err != nil {
			log.Fatal(err)
	}

	CreateUserTable(db)

	// 	{ // Create a new table
	// 		query := `
	// 			CREATE TABLE IF NOT EXISTS users (
	// 				id INT AUTO_INCREMENT,
	// 				username TEXT NOT NULL,
	// 				password TEXT NOT NULL,
	// 				created_at DATETIME,
	// 				PRIMARY KEY (id)
	// 			);`

	// 		if _, err := db.Exec(query); err != nil {
	// 				log.Fatal(err)
	// 		}
	// }

	{ // Insert a new user
			username := "johndoe"
			password := "secret"
			createdAt := time.Now()

			result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
			if err != nil {
					log.Fatal(err)
			}

			id, err := result.LastInsertId()
			fmt.Println(id)
	}

	{ // Query a single user
			var (
					id        int
					username  string
					password  string
					createdAt time.Time
			)

			query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
			if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
					log.Fatal(err)
			}

			fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
			type user struct {
					id        int
					username  string
					password  string
					createdAt time.Time
			}

			rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
			if err != nil {
					log.Fatal(err)
			}
			defer rows.Close()

			var users []user
			for rows.Next() {
					var u user

					err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
					if err != nil {
							log.Fatal(err)
					}
					users = append(users, u)
			}
			if err := rows.Err(); err != nil {
					log.Fatal(err)
			}

			fmt.Printf("%#v", users)
	}

	{
			_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
			if err != nil {
					log.Fatal(err)
			}
	}

	
	

	router := mux.NewRouter()
	router.HandleFunc("/", dashboard)
	router.HandleFunc("/books/{title}/page/{page}", bookController)


	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8001", router)
}

