package main

import (
	"os"

	"github.com/ribeirosaimon/skadi/api/internal/skadiEngine"
)

func main() {
	// f, err := os.Create("trace_referencia.out")
	// if err != nil {
	// 	log.Fatalf("Failed to create trace output file: %v", err)
	// }
	// defer f.Close()
	// if err := trace.Start(f); err != nil {
	// 	log.Fatalf("Failed to start trace: %v", err)
	// }
	skadiEngine.StartSkadiApi(os.Args[1])
}
