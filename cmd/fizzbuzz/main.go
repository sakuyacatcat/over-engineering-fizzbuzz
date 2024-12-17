package main

import (
	"fmt"
	"log"

	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/di"
)

func main() {
    // DIコンテナで必要な依存を構築
    app, err := di.InitializeApp()
    if err != nil {
        log.Fatalf("failed to initialize app: %v", err)
    }

    results, err := app.Run(1, 100)
    if err != nil {
        log.Fatalf("error running fizzbuzz: %v", err)
    }

    for _, r := range results {
        fmt.Println(r)
    }
}
