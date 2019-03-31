package surveys

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	cL "github.com/jjmarsha/NormsBackend/pkg/classes"
)

func SurveyHandler(db *sql.DB, u *cL.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var sV cL.Survey
		err1 := decoder.Decode(&sV)
		if err1 != nil {
			panic(err1)
		}

		surv := u.Uname + "survey"
		injection := "INSERT INTO " + surv + " (date, gluten, sugar, satfats, alcohol, refgrains, msg, salt) VALUES (?,?,?,?,?,?,?)"
		_, err := db.Query(injection, sV.Date, sV.Gluten, sV.Sugar, sV.Satfats, sV.Alcohol, sV.Refgrains, sV.Msg, sV.Salt)
		if err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			fmt.Println(err)
		}
	}
	return http.HandlerFunc(fn)
}

//MedSender sends medical info to app
// func SurveyDataSender(db *sql.DB, u *cL.User) http.HandlerFunc {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		md := cL.Med{}
// 		row := db.QueryRow("SELECT name, gender, race, age, weight, history FROM medhistory WHERE username=?", u.Uname)
// 		if err := row.Scan(&md.Name, &md.Gender, &md.Race, &md.Age, &md.Weight, &md.History); err != nil {
// 			panic(err.Error())
// 			return
// 		}
// 		js, err := json.Marshal(md)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(js)
// 	}
// 	return http.HandlerFunc(fn)
// }
