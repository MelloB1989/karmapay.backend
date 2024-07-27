package api

import (
	"karmapay/database"
	"log"

	_ "github.com/lib/pq"
)

func CreateAPIKey(API database.APIKeys){
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := db.Exec(`INSERT INTO api_keys (uid, id, api_key, pg_enum) VALUES ($1, $2, $3, $4)`, API.UID, API.ID, API.APIKey, API.PGEnum)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}

func ListAPIKeys(uid string) []database.APIKeys {
    db, err := database.DBConn()
    if err != nil {
        log.Fatalln(err)
    }

    rows, err := db.Query(`SELECT * FROM api_keys WHERE uid = $1`, uid)
    if err != nil {
        log.Fatalln(err)
    }
    defer rows.Close()

    var apiKeys []database.APIKeys
    for rows.Next() {
        var API database.APIKeys
        err := rows.Scan(&API.UID, &API.ID, &API.APIKey, &API.PGEnum)
        if err != nil {
            log.Fatalln(err)
        }
        apiKeys = append(apiKeys, API)
    }

    if err = rows.Err(); err != nil {
        log.Fatalln(err)
    }

    return apiKeys
}

func GetAPIKeyByUIDAndPGEnum (uid string, pgEnum string) database.APIKeys {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	var API database.APIKeys
	err = db.QueryRow(`SELECT * FROM api_keys WHERE uid = $1 AND pg_enum = $2`, uid, pgEnum).Scan(&API.UID, &API.ID, &API.APIKey, &API.PGEnum)
	if err != nil {
		log.Fatalln(err)
	}

	return API
}