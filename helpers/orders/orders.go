package users

import (
	"karmapay/database"
	"log"
	"github.com/redis/go-redis/v9"
	"karmapay/config"
	"encoding/json"
	"context"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

func CreateOrder(Order database.Order){
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := db.Exec(`INSERT INTO orders (uid, order_id, order_amount, order_currency, order_description, order_status, order_timestamp, order_upi_trnx_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, Order.UID, Order.OrderID, Order.OrderAmount, Order.OrderCurrency, Order.OrderDescription, Order.OrderStatus, Order.OrderTimeStamp, Order.OrderUpiTransactionID)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}

func PushOrderToRedis(Order database.RedisOrder) {
    opt, _ := redis.ParseURL(config.NewConfig().RedisURL)
    client := redis.NewClient(opt)

    // Stringify the RedisOrder struct
    orderJSON, err := json.Marshal(Order)
    if err != nil {
        log.Fatalln("Error stringifying order:", err)
    }

    client.Set(ctx, Order.OrderID, orderJSON, 0)
}