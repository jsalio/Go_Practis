# CSV Processor

A Go application for processing CSV files containing product inventory data, generating reports, and calculating inventory statistics.

## Description

This application reads a CSV file containing product information, processes the data, and generates a detailed report in a new CSV file. It calculates various statistics including total inventory value and identifies the most expensive product.

## Features

- CSV file reading and parsing
- Data validation and error handling
- Product inventory calculations
- Report generation with detailed statistics
- Automatic output file creation

## Input CSV Format

The input CSV file should have the following format:
```
ProductName,Price,Quantity
```

Example:
```
Laptop,999.99,5
Mouse,29.99,10
Keyboard,49.99,8
```

## Output

The program generates an `output.csv` file containing:
1. Header row with column names
2. Processed data with calculated values
3. Information about the most expensive product
4. Total inventory value

## Code Structure

### Main Components

1. **Data Structures**
   - `CsvData`: Represents raw product data from input
   - `CsvReport`: Represents processed data with calculated values

2. **Core Functions**
   - `ReadCsvFile`: Reads and parses the input CSV file
   - `ParceLine`: Processes individual CSV lines
   - `GenerateResume`: Creates the final report
   - `saveDataInReport`: Writes the report to output.csv

3. **Utility Functions**
   - `ParceNumericField`: Converts string to integer
   - `ParceFloatField`: Converts string to float
   - `FindMoreExpensiveItem`: Identifies the most expensive product

## Usage

1. Compile the program:
   ```bash
   go build main.go
   ```

2. Run the executable:
   ```bash
   ./main
   ```

3. When prompted, enter the name of your CSV file:
   ```
   Ingrese el archivo a leer (*.csv) :your_file.csv
   ```

4. The program will process the file and create `output.csv` in the same directory.

## Error Handling

The program includes comprehensive error handling for:
- File reading errors
- Invalid CSV formats
- Invalid numeric values
- Missing or malformed data

## Dependencies

- Standard Go libraries:
  - `encoding/csv`
  - `fmt`
  - `os`
  - `strconv`
  - `strings`

## Example

Input file (`products.csv`):
```
Product,Price,Quantity
Laptop,999.99,5
Mouse,29.99,10
Keyboard,49.99,8
```

Output file (`output.csv`):
```
Producto,Precio,Cantidad,Total
Laptop,999.99,5,4999.95
Mouse,29.99,10,299.90
Keyboard,49.99,8,399.92
Producto mas caro,Laptop,999.99
Total Inventario,5699.77
```

## Contributing

Feel free to submit issues and enhancement requests! 