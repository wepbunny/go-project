package excel

import (
	"fmt"
	"log"
	"sync"

	"excelproject/db"

	"github.com/xuri/excelize/v2"
)

// GenerateExcel generates the Excel file with multiple sheets.
func GenerateExcel() error {
	f := excelize.NewFile()

	var wg sync.WaitGroup
	wg.Add(3)

	// Generate summary sheet
	go func() {
		defer wg.Done()
		err := generateSummarySheet(f)
		if err != nil {
			log.Printf("Error generating summary sheet: %v", err)
		}
	}()

	// Generate high-amount and high-likes sheet
	go func() {
		defer wg.Done()
		err := generateHighLikesSheet(f)
		if err != nil {
			log.Printf("Error generating high-amount and high-likes sheet: %v", err)
		}
	}()

	// Generate low-amount and low-likes sheet
	go func() {
		defer wg.Done()
		err := generateLowLikesSheet(f)
		if err != nil {
			log.Printf("Error generating low-amount and low-likes sheet: %v", err)
		}
	}()

	wg.Wait()

	// Save the Excel file
	err := f.SaveAs("output.xlsx")
	if err != nil {
		return fmt.Errorf("failed to save Excel file: %v", err)
	}

	return nil
}

// generateSummarySheet generates the summary sheet with total orders and total sum amount.
func generateSummarySheet(f *excelize.File) error {
	orders, err := db.GetOrders()
	if err != nil {
		return err
	}

	// Set sheet name and headers
	sheetName := "Summary"
	f.SetSheetName("Sheet1", sheetName)
	f.SetCellValue(sheetName, "A1", "Total Orders")
	f.SetCellValue(sheetName, "B1", "Total Sum Amount")

	// Populating data
	totalOrders := len(orders)
	totalSumAmount := 0.0

	for _, order := range orders {
		totalSumAmount += order.Price * float64(order.Quantity)
	}

	f.SetCellValue(sheetName, "A2", totalOrders)
	f.SetCellValue(sheetName, "B2", totalSumAmount)

	return nil
}

// generateHighLikesSheet generates the sheet with orders having amount > 100 and likes > 10.
func generateHighLikesSheet(f *excelize.File) error {
	orders, err := db.GetHighLikesOrders()
	if err != nil {
		return err
	}

	// Set sheet name and headers
	sheetName := "High Likes"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "S.NO.")
	f.SetCellValue(sheetName, "B1", "ID")
	f.SetCellValue(sheetName, "C1", "Title")
	f.SetCellValue(sheetName, "D1", "Image")
	f.SetCellValue(sheetName, "E1", "Price")
	f.SetCellValue(sheetName, "F1", "Quantity")
	f.SetCellValue(sheetName, "G1", "Description")
	f.SetCellValue(sheetName, "H1", "Total Likes") // New column

	// Populate data
	row := 2
	for _, order := range orders {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), row-1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), order.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), order.Title)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), order.Image)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), order.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), order.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), order.Description)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), order.NumLikes)
		row++
	}

	return nil
}

// generateLowLikesSheet generates the sheet with orders having amount < 100 and likes < 3.
func generateLowLikesSheet(f *excelize.File) error {
	orders, err := db.GetLowLikesOrders()
	if err != nil {
		return err
	}

	// Set sheet name and headers
	sheetName := "Low Likes"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "S.NO.")
	f.SetCellValue(sheetName, "B1", "ID")
	f.SetCellValue(sheetName, "C1", "Title")
	f.SetCellValue(sheetName, "D1", "Image")
	f.SetCellValue(sheetName, "E1", "Price")
	f.SetCellValue(sheetName, "F1", "Quantity")
	f.SetCellValue(sheetName, "G1", "Description")
	f.SetCellValue(sheetName, "H1", "Total Likes") // New column

	// Populate data
	row := 2
	for _, order := range orders {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), row-1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), order.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), order.Title)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), order.Image)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), order.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), order.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), order.Description)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), order.NumLikes)
		row++
	}

	return nil
}
