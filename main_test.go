package main

import (
	"excelproject/db"
	"excelproject/excel"
	"testing"
)

func TestMain(t *testing.T) {
	// Initialize the database connection
	err := db.InitializeDB()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.CloseDB()

	// TODO: Add test cases for the main function
	// For example, you can simulate user input/output and check if the expected actions are performed.

	// Generate the Excel file
	err = excel.GenerateExcel()
	if err != nil {
		t.Fatalf("Failed to generate Excel file: %v", err)
	}

	// TODO: Add assertions to validate the output Excel file
	// For example, you can check if the file exists and has the expected data.
}
