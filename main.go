package main

import (
	"log"

	"excelproject/db"
	"excelproject/excel"
)

func main() {
	// Initialize the database connection
	err := db.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.CloseDB()

	// Generate the Excel file
	err = excel.GenerateExcel()
	if err != nil {
		log.Fatalf("Failed to generate Excel file: %v", err)
	}

	log.Println("Excel file generated successfully")
}
