package tests

import (
	"testing"

	"github.com/jakofys/paywall/internal/package/btc"
)

func TestCreateTransactionBTC(t *testing.T) {
	addr := ""
	amount := 1

	if err := btc.CreateLNTransactionBTC(addr, amount); err != nil {
		t.Errorf("RPC remote service BTC testnet not work for %s adresses", addr)
	}
}

func TestConnectionRPCServer(t *testing.T) {
	addr := ""

	if err := btc.PingRemoteBitcoinServer(addr); err != nil {
		t.Errorf("API connection to service BTC testnet failed, please verify configuration for %s", addr)
	}
}
