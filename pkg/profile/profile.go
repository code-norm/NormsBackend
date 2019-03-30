// package profile

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/jjmarsha/NormsBackend/pkg/session"
// )

// //VPHandler sends profile information form db
// func VPHandler(db *sql.DB, u *session.User) http.HandlerFunc {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		var viewPosts string

// 		rows, err := db.Query("SELECT 'one' col1, 'two' col2, 3 col3, NULL col4")
// 		if err != nil {
// 			fmt.Println("Failed to run query", err)
// 			return
// 		}

// 		cols, err := rows.Columns()
// 		if err != nil {
// 			fmt.Println("Failed to get columns", err)
// 			return
// 		}

// 		// Result is your slice string.
// 		rawResult := make([][]byte, len(cols))
// 		result := make([]string, len(cols))

// 		dest := make([]interface{}, len(cols)) // A temporary interface{} slice
// 		for i := range rawResult {
// 			dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
// 		}

// 		for rows.Next() {
// 			err = rows.Scan(dest...)
// 			if err != nil {
// 				fmt.Println("Failed to scan row", err)
// 				return
// 			}

// 			for i, raw := range rawResult {
// 				if raw == nil {
// 					result[i] = "\\N"
// 				} else {
// 					result[i] = string(raw)
// 				}
// 			}

// 			fmt.Printf("%#v\n", result)
// 		}
// 	}
// }
