package main

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load()

	if err != nil {
		t.Fatal(".env not getting loaded")
	}
}

func TestAppRunningOnEnvPort(t *testing.T) {
	t.Setenv("PORT", "8000")
	err := godotenv.Load()

	if err != nil {
		t.Error(".env file not loading")
	}

	port := os.Getenv("PORT")

	if port != "8000" {
		t.Error(".env port not recognised")
	}

}