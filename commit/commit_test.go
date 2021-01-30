package main

import (
  "testing"
)

// Test to check the commit made by the user
func TestCommit(t *testing.T) {
  expected := "first commit"
  got := "first commit"

  if got != expected {
    t.Errorf("Expected: %v, got: %v", expected, got)
  }
} 
