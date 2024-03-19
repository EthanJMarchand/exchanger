package main

import "testing"

func TestLoadEnv(t *testing.T) {
	config, err := loadEnv()
	if err != nil {

		t.Errorf("Loadenv() err: %s", err)
	}
	if config.CCKey == "" {
		t.Error("api key cannot be blank")
	}
	if config.URL == "" {
		t.Error("api url has not been set.")
	}
}
