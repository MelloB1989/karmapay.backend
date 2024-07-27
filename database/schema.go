package database

import "encoding/json"

type Customer struct {
	UID string `json:"uid"`
	CID string `json:"cid"`
	C_Email string `json:"c_email"`
	C_IP string `json:"c_ip"`
	C_Name string `json:"c_name"`
	C_Phone string `json:"c_phone"`
	C_Location string `json:"c_location"`
}

type Order struct {
	UID string `json:"uid"`
	OrderID string `json:"order_id"`
	OrderAmount string `json:"order_amount"`
	OrderStatus string `json:"order_status"`
	OrderCID string `json:"order_cid"`
	OrderCurrency string `json:"order_currency"`
	OrderDescription string `json:"order_description"`
	OrderTimeStamp string `json:"order_timestamp"`
	OrderUpiTransactionID string `json:"order_upi_transaction_id"`
}

type APIKeys struct {
	UID string `json:"uid"`
	ID string `json:"id"`
	APIKey string `json:"api_key"`
	PGEnum string `json:"pg_enum"`
}

type User struct {
	UID string `json:"uid"`
	Username string `json:"username"`
	BusinessName string `json:"business_name"`
	BusinessURL string `json:"business_url"`
	PFP string `json:"pfp"`
	Subdomain string `json:"subdomain"`
	Password string `json:"password"`
}

type RedisOrder struct {
	OrderID string `json:"order_id"`
	OrderStatus string `json:"order_status"`
	UID string `json:"uid"`
	Email string `json:"email"`
	KPAPI string `json:"kpapi"`
	API_KEY string `json:"api_key"`
	OrderAmount string `json:"order_amt"`
	OrderCurrency string `json:"order_currency"`
	OrderDescription string `json:"order_description"`
	Subdomain string `json:"subdomain"`
	OrderMode string `json:"order_mode"`
	WebhookURL string `json:"webhook_url"`
	RedirectURL string `json:"redirect_url"`
	Registration string `json:"registration"`
	OrderCID string `json:"order_cid"`
	PGOrder json.RawMessage `json:"PGOrder"`
	Timestamp string `json:"timestamp"`
}