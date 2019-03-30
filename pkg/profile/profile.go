package profile

import (
	"database/sql"
	"net/http"
	"NormsBackend/pkg/session"
)
//SymptomHandler receives the symptoms chosen to add to table
func SymptomHandler(db *sql.DB, u *User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		symptoms := make([]bool, 11)
		r.ParseForm()
		for i:=0; i < 11; i++ {
			symptoms[i] = r.FormValue(i);
		}

		if symptoms[0]
		{
			_, err = db.Query("UPDATE userinfo SET speech = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[1]
		{
			_, err = db.Query("UPDATE [table] SET vision = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[2]
		{
			_, err = db.Query("UPDATE [table] SET muscle = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[3]
		{
			_, err = db.Query("UPDATE [table] SET bladder = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[4]
		{
			_, err = db.Query("UPDATE [table] SET depression = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[5]
		{
			_, err = db.Query("UPDATE [table] SET memory = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[6]
		{
			_, err = db.Query("UPDATE [table] SET attention = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[7]
		{
			_, err = db.Query("UPDATE [table] SET mood = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[8]
		{
			_, err = db.Query("UPDATE [table] SET reasoning = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[9]
		{
			_, err = db.Query("UPDATE [table] SET dizziness = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
		if symptoms[10]
		{
			_, err = db.Query("UPDATE [table] SET judgement  = '1' WHERE username = ?",
				newUser.Uname)
			if err != nil {
				// If there is any issue with inserting into the database, return a 500 error
				fmt.Println(err)
				return
			}
		}
	}
}
