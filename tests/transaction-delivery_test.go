package tests

import "testing"

func TestCreateTransactionBTC(t *testing.T) {
	addr := ""
	amount := 1

	if err := CreateTransactionBTC(addr, amount); err != nil {
		t.Errorf("RPC remote service BTC testnet not work for %s adresses", addr)
	}
}

func TestConnectionRPCServer(t *testing.T) {
	addr := ""

	if ok, _ := PingRemoteBitcoinServer(addr); ok {
		t.Errorf("RPC remote connection to service BTC testnet failed, please verify configuration", addr)
	}
}
