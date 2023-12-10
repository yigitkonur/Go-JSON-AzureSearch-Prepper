package main

import (
    "fmt"
    "log"
)

func main() {
    config := ProcessorConfig{
        JSONInputPath:  "/path/to/json/input",
        JSONOutputPath: "/path/to/json/output.json",
        CSVOutputPath:  "/path/to/csv/output.csv",
    }

    fileProcessor := NewFileProcessor(config)

    if err := fileProcessor.StartProcessing(); err != nil {
        log.Fatalf("Error processing files: %v", err)
    }

    fmt.Println("Processing complete.")
}
