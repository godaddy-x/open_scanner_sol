package handlers

import (
	"github.com/godaddy-x/freego/rpcx"
)

// Register adds Solana-specific RPC handlers.
// Core create/submit/summary/scan APIs live in open_scanner generic walletapi;
// Solana v1 has no SpeedUp/Cancel/batch/deploy extensions yet.
func Register(m *rpcx.RPCManager) {
	_ = m
}
