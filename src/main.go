package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jjmarsha/NormsBackend/pkg/profile"
	"github.com/jjmarsha/NormsBackend/pkg/session"
)

func main() {

	mux := http.NewServeMux()

	db, err := sql.Open("mysql", "root:root@tcp(35.197.6.145:3306)/userinfo")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping errored: %v", err)
		panic(err.Error()) // proper error handling instead of panic
	}

	CurrUser := session.User{
		Uname:  "**noCurr**",
		Uemail: "**noCurr**",
	}

	// // Signup
	// mux.HandleFunc("/signup", session.SignupHandler(db, &CurrUser))

	// //Login
	// mux.HandleFunc("/login", session.LoginHandler(db, &CurrUser))

	// //Gets the post. Will relocate
	// mux.HandleFunc("/viewpost", post.Handler(db, &CurrUser))
	mux.HandleFunc("/symptoms", profile.SymptomHandler(db, &CurrUser))

	fmt.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
