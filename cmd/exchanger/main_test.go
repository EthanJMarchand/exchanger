package main

import (
	"testing"
)

// func Test_loadEnv(t *testing.T) {
// 	type args struct {
// 		filenames []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Config
// 		wantErr bool
// 	}{
// 		{"testing error", args{filenames: "./testdata/.env.test"},}, // TODO: get unstuck lol.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := loadEnv(tt.args.filenames...)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("loadEnv() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("loadEnv() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestLoadEnv(t *testing.T) {
	config, err := loadEnv("./testdata/.env.test")
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
