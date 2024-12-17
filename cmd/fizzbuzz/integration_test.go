package main

import (
	"testing"

	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/di"
)

func TestFizzBuzzIntegration(t *testing.T) {
    app, err := di.InitializeApp()
    if err != nil {
        t.Fatalf("failed to initialize app: %v", err)
    }

    results, err := app.Run(1, 15)
    if err != nil {
        t.Fatalf("error running fizzbuzz: %v", err)
    }

    expected := []string{
        "1", "2", "Fizz", "4", "Buzz",
        "Fizz", "7", "8", "Fizz", "Buzz",
        "11", "Fizz", "13", "14", "FizzBuzz",
    }

    if len(results) != len(expected) {
        t.Fatalf("expected length %d, got %d", len(expected), len(results))
    }

    for i, v := range results {
        if v != expected[i] {
            t.Errorf("at %d expected %q, got %q", i, expected[i], v)
        }
    }
}
