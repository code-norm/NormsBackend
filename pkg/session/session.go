package session

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	cL "github.com/jjmarsha/NormsBackend/pkg/classes"
	"golang.org/x/crypto/bcrypt"
)

//IsEmpty will check for empty entries
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

//SignupHandler handles signup
func SignupHandler(db *sql.DB, u *cL.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		newUser := cL.User{
			Uname:  r.FormValue("username"), // Data from the form
			Uemail: r.FormValue("email"),    // Data from the form
		}

		fmt.Println("signup request parsed")

		Upwd := r.FormValue("password")
		pwdConfirm := r.FormValue("confirm") // Data from the form

		// Empty data checking
		uNameCheck := IsEmpty(newUser.Uname)
		emailCheck := IsEmpty(newUser.Uemail)
		pwdCheck := IsEmpty(Upwd)
		pwdConfirmCheck := IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}

		if Upwd != pwdConfirm {
			// Save to database (username, email and password)
			fmt.Printf("Passwords are not matching")
			return
		}

		// Checks for existing username
		// if err1 := db.QueryRow("SELECT * FROM userinfo WHERE username=?", newUser.Uname); err1 == nil {
		// 	fmt.Fprintf(w, "username already exists")
		// 	return
		// }

		// // Checks for existing email
		// if err2 := db.QueryRow("SELECT * FROM userinfo WHERE email=?", newUser.Uemail); err2 == nil {
		// 	fmt.Fprintf(w, "Email already in used")
		// 	return
		// }
		if rowExists("SELECT * FROM userinfo WHERE username=?", db, newUser.Uname) {
			fmt.Fprintf(w, "Username in use")
			return
		}

		if rowExists("SELECT * FROM userinfo WHERE email=?", db, newUser.Uemail) {
			fmt.Fprintf(w, "Email in use")
			return
		}

		fmt.Println("signup request confirmed")

		encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(Upwd), 8)
		fmt.Println("password encrypted")

		// Inserts new record
		_, err = db.Query("INSERT INTO userinfo (username, password, email) VALUES(?, ?, ?)",
			newUser.Uname, string(encryptedPwd), newUser.Uemail)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}

		// Inserts notifs
		_, err = db.Query("INSERT INTO notif (username) VALUES(?)", newUser.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}

		// Inserts medhistory
		_, err = db.Query("INSERT INTO medhistory (username) VALUES(?)", newUser.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		*u = newUser

		surv := u.Uname + "survey"
		injection := "CREATE TABLE " + surv + " (date VARCHAR(12), gluten VARCHAR(60), sugar VARCHAR(60), satfat VARCHAR(60), alcohol VARCHAR(60), refgrains VARCHAR(60), msg VARCHAR(60), salt VARCHAR(60))"
		// Creates new survey entry for new users
		_, err = db.Query(injection)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}

		js, err := json.Marshal(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	return http.HandlerFunc(fn)
}

//LoginHandler handles login request
func LoginHandler(db *sql.DB, u *cL.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email := r.FormValue("email")  // Data from the form
		pwd := r.FormValue("password") // Data from the form

		// Empty data checking
		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}
		var encryptedPwd string
		var username string

		// Searches for email and grabs encrypted password
		row := db.QueryRow("SELECT password,username FROM userinfo WHERE email=?", email)
		switch err := row.Scan(&encryptedPwd, &username); err {
		case sql.ErrNoRows:
			fmt.Fprintf(w, "No users with email")
			return
		case nil:
			fmt.Println("user found")
		default:
			fmt.Println(err)
			return
		}

		// Compares encrypted password to the entered password
		err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(pwd))
		if err != nil {
			fmt.Fprintf(w, "wrong password")
			return
		}

		u.Uemail = email
		u.Uname = username

		userTemp := cL.Usertemp{
			Username: username,
			Email:    email,
		}

		js, err := json.Marshal(userTemp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	return http.HandlerFunc(fn)
}

func LogoutHandler(u *cL.User) http.HandlerFunc {
	u.Uname = "NULL"
	u.Uemail = "NULL"
	fn := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.HandlerFunc(fn)
}

func rowExists(query string, db *sql.DB, entry string) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, entry).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("error checking if row exist")
	}
	return exists
}

func test(r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		fmt.Printf("%s = s%\n", key, value)
	}
}
