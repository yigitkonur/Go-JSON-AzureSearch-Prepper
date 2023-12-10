# Go-JSON-AzureSearch-Prepper
A robust and efficient Go utility for processing large volumes of JSON files, preparing them for Azure Search AI. This tool handles concurrent processing of multiple JSON files, compiles them into a singular JSON and CSV format, and assigns unique identifiers to each entry, making the data ready for further use in Azure Search AI.

## Features
- **Concurrent JSON File Processing**: Optimized for handling large datasets with speed and efficiency.
- **Unique Identifier Assignment**: Each JSON entry is assigned a unique identifier, ensuring data integrity.
- **Customizable Output**: Generates both JSON and CSV outputs, suitable for various use cases including Azure Search AI.
- **Progress Tracking**: Includes a progress bar for real-time processing updates.
- **Modular Design**: Easily extendable for additional data processing needs.

## Use Cases
- **Azure Search AI Preparation**: Prepare and aggregate data from multiple JSON files for Azure Search AI.
- **Data Transformation**: Transform JSON data into a structured CSV format for analytics and reporting.
- **Data Integration**: Integrate and consolidate disparate JSON data sources for unified processing and analysis.

## Getting Started

### Prerequisites
- Go (version 1.16 or later)
- Basic understanding of Go project structure and modules

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/Go-JSON-AzureSearch-Prepper.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Go-JSON-AzureSearch-Prepper
   ```
3. Install dependencies:
   ```bash
   go get -v ./...
   ```

### Usage
1. Update the `processorConfig` in `main.go` with the paths to your input JSON files, output JSON file, and output CSV file.
2. Run the program:
   ```bash
   go run main.go
   ```
