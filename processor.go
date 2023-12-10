package processor

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "sync/atomic"

    "github.com/schollz/progressbar/v3"
    jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var currentID int64

// FileProcessor is an implementation of DataProcessor for processing files.
type FileProcessor struct {
    Config ProcessorConfig
    wg     sync.WaitGroup
    jsonCh chan JSONEntry
    csvCh  chan []string
}

// NewFileProcessor creates a new FileProcessor with the given configuration.
func NewFileProcessor(config ProcessorConfig) *FileProcessor {
    return &FileProcessor{
        Config: config,
        jsonCh: make(chan JSONEntry, 100),
        csvCh:  make(chan []string, 100),
    }
}

// StartProcessing starts the file processing.
func (fp *FileProcessor) StartProcessing() error {
    files, err := os.ReadDir(fp.Config.JSONInputPath)
    if err != nil {
        return err
    }

    bar := progressbar.Default(int64(len(files)))

    go fp.writeJSON(fp.Config.JSONOutputPath)
    go fp.writeCSV(fp.Config.CSVOutputPath)

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".json" {
            fp.wg.Add(1)
            go fp.processFile(filepath.Join(fp.Config.JSONInputPath, file.Name()), bar)
        }
    }

    fp.wg.Wait()
    close(fp.jsonCh)
    close(fp.csvCh)

    return nil
}

func (fp *FileProcessor) processFile(filePath string, bar *progressbar.ProgressBar) {
    defer fp.wg.Done()

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        var entry map[string]interface{}
        err := json.Unmarshal(scanner.Bytes(), &entry)
        if err != nil {
            fmt.Println("Error unmarshalling JSON:", err)
            continue
        }

        id := atomic.AddInt64(&currentID, 1)

        fp.jsonCh <- JSONEntry{ID: fmt.Sprint(id), Document: entry["document"].(string)}
        fp.csvCh <- []string{fmt.Sprint(id), entry["country"].(string), entry["language"].(string), entry["keyword"].(string), fmt.Sprint(entry["search_volume"].(float64))}
    }

    bar.Add(1)
}

func (fp *FileProcessor) writeJSON(outputPath string) {
    file, err := os.Create(outputPath)
    if err != nil {
        fmt.Println("Error creating JSON file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    for entry := range fp.jsonCh {
        data, err := json.Marshal(entry)
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            continue
        }
        writer.Write(data)
        writer.WriteString("\n")
    }
}

func (fp *FileProcessor) writeCSV(outputPath string) {
    file, err := os.Create(outputPath)
    if err != nil {
        fmt.Println("Error creating CSV file:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(bufio.NewWriter(file))
    defer writer.Flush()

    for record := range fp.csvCh {
        if err := writer.Write(record); err != nil {
            fmt.Println("Error writing record to CSV:", err)
        }
    }
}
