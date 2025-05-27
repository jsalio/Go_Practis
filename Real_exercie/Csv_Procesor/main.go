// Package main provides functionality to process CSV files containing product data,
// generate reports, and calculate inventory statistics.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CsvData represents a single row of product data from the CSV file.
type CsvData struct {
	product_name     string  // Name of the product
	product_price    float64 // Price of the product
	product_quantity int     // Quantity of the product in stock
}

// CsvReport represents a processed row of data with calculated total value.
type CsvReport struct {
	product_name     string  // Name of the product
	product_price    float64 // Price of the product
	product_quantity int     // Quantity of the product in stock
	product_value    float64 // Total value (price * quantity)
}

// ParceNumericField converts a string field to an integer.
// It trims whitespace and validates the conversion.
// Returns an error if the field cannot be converted to an integer.
func ParceNumericField(field string) (int, error) {
	field = strings.TrimSpace(field)
	num, err := strconv.Atoi(field)
	if err != nil {
		return 0, fmt.Errorf("invalid numeric field: %s", field)
	}
	return num, nil
}

// ParceFloatField converts a string field to a float64.
// It trims whitespace and validates the conversion.
// Returns an error if the field cannot be converted to a float.
func ParceFloatField(field string) (float64, error) {
	field = strings.TrimSpace(field)
	num, err := strconv.ParseFloat(field, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float field: %s", field)
	}
	return num, nil
}

// ParceLine processes a single line from the CSV file and converts it to a CsvData struct.
// It expects at least 3 fields: product name, price, and quantity.
// Returns an error if the line format is invalid or if numeric fields cannot be parsed.
func ParceLine(line []string) (CsvData, error) {
	if len(line) < 3 {
		return CsvData{}, fmt.Errorf("invalid line: %s", line)
	}
	productName := line[0]
	productPrice, err := ParceFloatField(line[1])
	if err != nil {
		return CsvData{}, fmt.Errorf("invalid price field: %s", line[1])
	}
	productQuantity, err := ParceNumericField(line[2])
	if err != nil {
		return CsvData{}, fmt.Errorf("invalid quantity field: %s", line[2])
	}
	return CsvData{
		product_name:     productName,
		product_price:    productPrice,
		product_quantity: productQuantity,
	}, nil
}

// ReadCsvFile reads and parses a CSV file containing product data.
// It expects the file to have at least 3 columns: product name, price, and quantity.
// Returns a slice of CsvData and any error encountered during reading or parsing.
func ReadCsvFile(filename string) ([]CsvData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []CsvData
	for _, record := range records {
		if len(record) < 3 {
			continue
		}
		csvData, err := ParceLine(record)
		if err != nil {
			continue
		}
		data = append(data, csvData)
	}
	return data, nil
}

// FindMoreExpensiveItem finds the product with the highest price in the dataset.
// Returns a CsvData struct containing the most expensive product's information.
func FindMoreExpensiveItem(data []CsvData) CsvData {
	var mostEpensive CsvData = CsvData{
		product_name:     "",
		product_price:    0,
		product_quantity: 0,
	}
	for _, row := range data {
		if row.product_price > mostEpensive.product_price {
			mostEpensive = row
		}
	}
	return mostEpensive
}

// GenerateResume processes the CSV data and generates a summary report.
// Returns:
//   - Header row for the report
//   - Processed data with calculated values
//   - Information about the most expensive product
//   - Total inventory value
func GenerateResume(data []CsvData) ([]string, []CsvReport, []string, []string) {
	fisrtLine := []string{"Producto", "Precio", "Cantidad", "Total"}
	var totalInventory float64 = 0
	var report []CsvReport
	for _, row := range data {
		report = append(report, CsvReport{
			product_name:     row.product_name,
			product_price:    row.product_price,
			product_quantity: row.product_quantity,
			product_value:    row.product_price * float64(row.product_quantity),
		})
		totalInventory += row.product_price * float64(row.product_quantity)
	}
	moreExpensive := FindMoreExpensiveItem(data)
	mostCost := []string{"Producto mas caro", moreExpensive.product_name, strconv.FormatFloat(moreExpensive.product_price, 'f', 2, 64)}
	InventoryTotal := []string{"Total Inventario", strconv.FormatFloat(totalInventory, 'f', 2, 64)}

	return fisrtLine, report, mostCost, InventoryTotal
}

// saveDataInReport writes the processed data to a new CSV file named "output.csv".
// It includes the header, processed data rows, and summary information.
// The file is created in the current directory.
func saveDataInReport(header []string, data []CsvReport, footer []string, resume []string) {
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(header)
	for _, row := range data {
		writer.Write([]string{row.product_name, strconv.FormatFloat(row.product_price, 'f', 2, 64), strconv.FormatFloat(float64(row.product_quantity), 'f', 2, 64), strconv.FormatFloat(row.product_value, 'f', 2, 64)})
	}
	writer.Write(footer)
	writer.Write(resume)

	if err := writer.Error(); err != nil {
		fmt.Println("Error al escribir en el CSV:", err)
	}
}

// main is the entry point of the program.
// It prompts the user for a CSV file name, processes the file,
// generates a report, and saves the results to output.csv.
func main() {
	fmt.Println("Csv file processor")
	file := ""
	fmt.Print("Ingrese el archivo a leer (*.csv) :")
	_, error := fmt.Scanln(&file)
	if error != nil {
		fmt.Println("Error reading file name")
		return
	}
	data, err := ReadCsvFile(file)
	if err != nil {
		fmt.Println("Error reading file")
	}
	fmt.Println("Total products:", len(data))
	header, dataSet, footer, resume := GenerateResume(data)
	saveDataInReport(header, dataSet, footer, resume)
}
