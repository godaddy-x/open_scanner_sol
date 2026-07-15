package plugin

import (
	"strings"

	"github.com/godaddy-x/open_scanner/rpc/impl"
	"github.com/godaddy-x/open_scanner/scanner"
	"github.com/godaddy-x/open_scanner_sol/handlers"
	"github.com/godaddy-x/wallet-adapter-sol/sol"
)

func init() {
	sol.RegisterSignVerify()
	impl.RegisterHandlers(handlers.Register)
	// SOL 主币/SPL 为 address 模型、单笔 1:1，按地址精细判定流水（与 ETH/TRX 一致）。
	scanner.RegisterPerAddressFlowChecker(func(mainSymbol string) bool {
		switch strings.ToUpper(strings.TrimSpace(mainSymbol)) {
		case "SOL":
			return true
		default:
			return false
		}
	})
}
