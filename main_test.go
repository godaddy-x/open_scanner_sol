package main

import (
	"testing"

	"github.com/godaddy-x/open_scanner/run"
)

// TestRunSOL requires resource/config_ops.yaml and Mongo; skip by default in CI.
func TestRunSOL(t *testing.T) {
	*c = "resource/config_ops.yaml"
	run.Adapter(newAdapter, *c, "scanner_main", "SOL", "SOL", 0)
}
