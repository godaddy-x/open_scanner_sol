package main

import (
	"flag"
	"fmt"

	"github.com/godaddy-x/open_scanner/model"
	"github.com/godaddy-x/open_scanner/run"
	adapter "github.com/godaddy-x/wallet-adapter"
	"github.com/godaddy-x/wallet-adapter-sol/sol"

	_ "github.com/godaddy-x/open_scanner_sol/plugin"
)

var (
	name = flag.String("name", "ops-scanner-sol", "process name")
	c    = flag.String("config", "resource/config_ops.yaml", "config file path")
	m    = flag.String("mainsymbol", "SOL", "main symbol")
	s    = flag.String("symbol", "SOL", "symbol")
	i    = flag.Int64("initheight", 0, "init slot height")
)

func newAdapter(data *model.OwSymbol) adapter.ChainAdapter {
	decimals := int32(data.Decimals)
	if decimals <= 0 {
		decimals = 9
	}
	return sol.NewSolAdapter(data.Symbol, data.Name, decimals)
}

func main() {
	flag.Parse()
	fmt.Println(*name)
	fmt.Println("adapter symbol:", *s)
	run.Adapter(newAdapter, *c, "scanner_main", *m, *s, *i)
}
