package main

import (
	"os"
	"testing"
)

func TestOpenOrCreateFile(t *testing.T) {
	fileName := "test_creating_file"
	_ = openOrCreateFile(fileName)
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Error("File is not created")
	}
	// should raise panic if cannot open
	_ = openOrCreateFile(fileName)
	os.Remove(fileName)
}
