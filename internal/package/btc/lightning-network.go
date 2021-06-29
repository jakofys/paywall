package btc

import (
	"encoding/json"
	"net/http"
)

type Invoice struct {
	amount          int    `json:amount`
	description     string `json:description`
	expireIn        string `json:expireIn`
	addr            string `json:fallbackAddress`
	paymentPreimage string `json:paymentPreimage`
}

func CreateLNInvoiceBTC(addr string, msatAmount int) error {
	read := json.Marshal(Invoice{
		amount:          1,
		expireIn:        30,
		addr:            addr,
		paymentPreimage: 3q0,
	})
	http.Post("http://localhost:8080/createinvoice", "application:json", read)
	return nil
}

func PingRemoteBitcoinServer(addr string) error {
	return nil
}
