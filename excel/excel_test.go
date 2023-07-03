package excel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestGenerateSummarySheet(t *testing.T) {
	// Generate the Excel file
	err := GenerateExcel()
	assert.NoError(t, err)

	// Open the generated Excel file
	f, err := excelize.OpenFile("output.xlsx")
	assert.NoError(t, err)

	// Assert the content of the summary sheet
	_, aval := f.GetCellValue("Summary", "A1")
	_, bval := f.GetCellValue("Summary", "B1")
	assert.Equal(t, "Total Orders", aval)
	assert.Equal(t, "Total Sum Amount", bval)
	// Add more assertions to validate the content of the summary sheet

	// Close the Excel file
	err = f.Close()
	assert.NoError(t, err)

	// Clean up the output file
	err = os.Remove("output.xlsx")
	assert.NoError(t, err)
}

// func TestGenerateHighLikesSheet(t *testing.T) {
// 	// Create a new Excel file
// 	f := NewFile()

// 	// Generate the high-likes sheet
// 	err := generateHighLikesSheet(f)
// 	assert.NoError(t, err)

// 	// Validate the content of the high-likes sheet
// 	assert.Equal(t, "S.NO.", f.GetCellValue("High Likes", "A1"))
// 	assert.Equal(t, "ID", f.GetCellValue("High Likes", "B1"))
// 	assert.Equal(t, "Title", f.GetCellValue("High Likes", "C1"))
// 	assert.Equal(t, "Image", f.GetCellValue("High Likes", "D1"))
// 	assert.Equal(t, "Price", f.GetCellValue("High Likes", "E1"))
// 	assert.Equal(t, "Quantity", f.GetCellValue("High Likes", "F1"))
// 	assert.Equal(t, "Description", f.GetCellValue("High Likes", "G1"))
// 	assert.Equal(t, "Total Likes", f.GetCellValue("High Likes", "H1"))
// 	// Add more assertions to validate the content of the high-likes sheet
// }

// func TestGenerateLowLikesSheet(t *testing.T) {
// 	// Create a new Excel file
// 	f := NewFile()

// 	// Generate the low-likes sheet
// 	err := generateLowLikesSheet(f)
// 	assert.NoError(t, err)

// 	// Validate the content of the low-likes sheet
// 	assert.Equal(t, "S.NO.", f.GetCellValue("Low Likes", "A1"))
// 	assert.Equal(t, "ID", f.GetCellValue("Low Likes", "B1"))
// 	assert.Equal(t, "Title", f.GetCellValue("Low Likes", "C1"))
// 	assert.Equal(t, "Image", f.GetCellValue("Low Likes", "D1"))
// 	assert.Equal(t, "Price", f.GetCellValue("Low Likes", "E1"))
// 	assert.Equal(t, "Quantity", f.GetCellValue("Low Likes", "F1"))
// 	assert.Equal(t, "Description", f.GetCellValue("Low Likes", "G1"))
// 	assert.Equal(t, "Total Likes", f.GetCellValue("Low Likes", "H1"))
// 	// Add more assertions to validate the content of the low-likes sheet
// }

// func NewFile() *excelize.File {
// 	// Create a new Excel file
// 	f := excelize.NewFile()

// 	// Set the default sheet
// 	f.SetSheetName("Sheet1", "Summary")

// 	return f
// }

// func TestGenerateExcel(t *testing.T) {
// 	// Generate the Excel file
// 	err := GenerateExcel()
// 	assert.NoError(t, err)

// 	// Validate the existence of the output file
// 	_, err = os.Stat("output.xlsx")
// 	assert.False(t, os.IsNotExist(err))

// 	// Clean up the output file
// 	err = os.Remove("output.xlsx")
// 	assert.NoError(t, err)
// }
