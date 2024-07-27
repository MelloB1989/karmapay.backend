package customers

import (
	"karmapay/database"
	"log"

	_ "github.com/lib/pq"
)

func CreateCustomer(Customer database.Customer){
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := db.Exec(`INSERT INTO customers (uid, cid, c_email, c_location, c_name, c_phone) VALUES ($1, $2, $3, $4, $5, $6)`, Customer.UID, Customer.CID, Customer.C_Email, Customer.C_Location, Customer.C_Name, Customer.C_Phone)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}