package users

import (
	"karmapay/database"
	"log"

	_ "github.com/lib/pq"
)

func CreateUser(User database.User){
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := db.Exec(`INSERT INTO users (uid, username, business_name, business_url, pfp, subdomain, password) VALUES ($1, $2, $3, $4, $5, $6, $7)`, User.UID, User.Username, User.BusinessName, User.BusinessURL, User.PFP, User.Subdomain, User.Password)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}