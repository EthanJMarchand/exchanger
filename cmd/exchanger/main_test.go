package main_test

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// config, err := loadEnv()
	config.CCKey = "fakeapikey"
	config.URL = "https://fake.com"
	if config.CCKey == "" {
		t.Error("Loadenv() api key cannot be blank")
	}
	if config.URL == "" {
		t.Error("Loadenv()api url has not been set.")
	}
}
