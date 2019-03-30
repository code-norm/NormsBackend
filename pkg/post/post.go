package post

import (
	"database/sql"
	"net/http"

	"github.com/jjmarsha/NormsBackend/pkg/session"
	"github.com/jjmarsha/NormsBackend/pkg/stack"
)

//UserPost struct to store a post
type UserPost struct {
	user     string
	date     string
	postText string
	comments stack.ItemStack
	likes    []string
}

//Handler will return the post and its info
func Handler(db *sql.DB, u *session.User) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
	}
}
