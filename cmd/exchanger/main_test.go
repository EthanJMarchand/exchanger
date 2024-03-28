package main

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
	config, err := loadEnv()
	if err != nil {
		t.Fatal("could not load env file")
	}
	if config.CCKey == "" {
		t.Error("Loadenv() api key cannot be blank")
	}
	if config.URL == "" {
		t.Error("Loadenv()api url has not been set.")
	}
}
