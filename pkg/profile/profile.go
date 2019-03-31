package profile

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jjmarsha/NormsBackend/pkg/session"
)

//SymptomHandler receives the symptoms chosen to add to table
func SymptomHandler(db *sql.DB, u *session.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var symptoms [12]session.Symptom
		err1 := decoder.Decode(&symptoms)
		if err1 != nil {
			panic(err1)
		}
		u.Symp = symptoms

		_, err := db.Query("UPDATE userinfo SET speech = ? WHERE username = ?", checkTrue(symptoms[0]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET vision = ? WHERE username = ?", checkTrue(symptoms[1]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET muscle = ? WHERE username = ?", checkTrue(symptoms[2]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET bladder = ? WHERE username = ?", checkTrue(symptoms[3]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET depression = ? WHERE username = ?", checkTrue(symptoms[4]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET memory = ? WHERE username = ?", checkTrue(symptoms[5]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET attention = ? WHERE username = ?", checkTrue(symptoms[6]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET mood = ? WHERE username = ?", checkTrue(symptoms[7]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET reasoning = ? WHERE username = ?", checkTrue(symptoms[8]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET dizziness = ? WHERE username = ?", checkTrue(symptoms[9]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE userinfo SET judgement  = ? WHERE username = ?", checkTrue(symptoms[10]),
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////

		_, err = db.Query("UPDATE notif SET speech = ? WHERE username = ?", symptoms[0].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET vision = ? WHERE username = ?", symptoms[1].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET muscle = ? WHERE username = ?", symptoms[2].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET bladder = ? WHERE username = ?", symptoms[3].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET depression = ? WHERE username = ?", symptoms[4].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET memory = ? WHERE username = ?", symptoms[5].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET attention = ? WHERE username = ?", symptoms[6].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET mood = ? WHERE username = ?", symptoms[7].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET reasoning = ? WHERE username = ?", symptoms[8].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET dizziness = ? WHERE username = ?", symptoms[9].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
		_, err = db.Query("UPDATE notif SET judgement  = ? WHERE username = ?", symptoms[10].Notifications,
			u.Uname)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func SymptomSender(db *sql.DB, u *session.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		js, err := json.Marshal(u.Symp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
	return http.HandlerFunc(fn)
}

func checkTrue(s session.Symptom) int {
	if s.Checked == "true" {
		return 1
	}
	return 0
}
