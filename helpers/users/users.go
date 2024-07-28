package users

import (
	"karmapay/database"
	"log"

	_ "github.com/lib/pq"
)

func CreateUser(User database.User){
	db, err := database.DBConn()
	if err != nil {
		//log.Fatalln(err)
		log.Println(err)
	}

	r, err := db.Exec(`INSERT INTO users (uid, username, business_name, business_url, pfp, subdomain, password) VALUES ($1, $2, $3, $4, $5, $6, $7)`, User.UID, User.Username, User.BusinessName, User.BusinessURL, User.PFP, User.Subdomain, User.Password)

	if err != nil || r == nil {
		//log.Fatalln(err)
		log.Println(err)
	}
}

func GetUserByUID(uid string) database.User {
	db, err := database.DBConn()
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}

	var User database.User
	err = db.QueryRow(`SELECT * FROM users WHERE uid = $1`, uid).Scan(&User.UID, &User.Username, &User.BusinessName, &User.BusinessURL, &User.PFP, &User.Subdomain, &User.Password)
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}

	return User
}

func GetUserByUsername(username string) database.User {
	db, err := database.DBConn()
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}

	var User database.User
	err = db.QueryRow(`SELECT * FROM users WHERE username = $1`, username).Scan(&User.UID, &User.Username, &User.Password, &User.BusinessName, &User.BusinessURL, &User.PFP, &User.Subdomain)
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}

	return User
}