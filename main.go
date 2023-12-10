package main

import (
    "fmt"
    "log"

    "github.com/yourusername/yourproject/processor"
)

func main() {
    config := processor.ProcessorConfig{
        JSONInputPath:  "/path/to/json/input",
        JSONOutputPath: "/path/to/json/output.json",
        CSVOutputPath:  "/path/to/csv/output.csv",
    }

    fileProcessor := processor.NewFileProcessor(config)

    if err := fileProcessor.StartProcessing(); err != nil {
        log.Fatalf("Error processing files: %v", err)
    }

    fmt.Println("Processing complete.")
}
