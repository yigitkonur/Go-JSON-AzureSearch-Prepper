package main

// JSONEntry defines the structure of the JSON data to be processed.
type JSONEntry struct {
    ID       string `json:"id"`
    Document string `json:"text"`
}

// ProcessorConfig defines configuration options for the file processor.
type ProcessorConfig struct {
    JSONInputPath  string
    JSONOutputPath string
    CSVOutputPath  string
}
