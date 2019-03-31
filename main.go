package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jjmarsha/NormsBackend/pkg/profile"
	"github.com/jjmarsha/NormsBackend/pkg/session"
)

func main() {

	router := mux.NewRouter()

	fmt.Println("Starting server...")

	db, err := sql.Open("mysql", "root:root@tcp(35.202.9.105:3306)/users")
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

	fmt.Println("Server connected")

	CurrUser := session.User{
		Uname:  "NULL",
		Uemail: "NULL",
	}

	// Signup
	router.HandleFunc("/signup", session.SignupHandler(db, &CurrUser))

	//Login
	router.HandleFunc("/login", session.LoginHandler(db, &CurrUser))

	router.HandleFunc("/logout", session.LogoutHandler(&CurrUser))

	//Push user's symptoms onto the struct and database
	router.HandleFunc("/postsymptoms", profile.SymptomHandler(db, &CurrUser))

	//Get user's symptoms
	router.HandleFunc("/getsymptoms", profile.SymptomSender(db, &CurrUser))

	//Adds medical records
	router.HandleFunc("/postmed", profile.MedHandler(db, &CurrUser))

	//Gets medical records
	router.HandleFunc("/getmed", profile.MedSender(db, &CurrUser))

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Starting server on port " + port)

	err = http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
	return
}
