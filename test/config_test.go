package test

import (
	"fmt"
	"kulkasku/internal/config"
	"testing"
)

func TestConfig(t *testing.T) {

	cfg, err := config.LoadConfigDatabase()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cfg)
}