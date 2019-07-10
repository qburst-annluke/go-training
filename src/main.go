package main

import (
  "database/sql"
  "fmt"
	"github.com/tealeg/xlsx"
  "./connect"
  "./config"
)

func main() {
	db := connect.ConnectDb()

  defer db.Close()

  excelToDb(db)
  dbToExcel(db)
}

func excelToDb(db *sql.DB) {
  sqlStatement := `INSERT INTO users (name, age, email) VALUES ($1, $2, $3)`

  xlFile, err := xlsx.OpenFile(config.SourceFile)
  if err != nil {
    fmt.Println("Error")
  }

  for _, sheet := range xlFile.Sheets {
    for rowCount, row := range sheet.Rows {
      if rowCount == 0 {
        continue
      }

      currentRow := make([]string, 3)
      for i, cell := range row.Cells {
        text := cell.String()
        currentRow[i] = text
      }
      _, err = db.Exec(sqlStatement, currentRow[0], currentRow[1], currentRow[2])
      if err != nil {
        fmt.Println("Error:", err)
      } else {
        fmt.Printf("Row %d added successfully \n", rowCount)
      }
    }
  }
}

func dbToExcel(db *sql.DB) {
	sqlStatement := `SELECT * FROM users`

  users, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	file := xlsx.NewFile()
  sheet, err := file.AddSheet(config.SheetName)
  if err != nil {
    fmt.Printf(err.Error())
  }

	for users.Next() {
		details := make([]string, 4)
		err := users.Scan(&details[0], &details[1], &details[2], &details[3])
		if err != nil {
			fmt.Println("Error: ", err)
		}
		row := sheet.AddRow()
		for _, value := range details {
			cell := row.AddCell()
			cell.Value = value
		}
		err = file.Save(config.DestinationFile)
  	if err != nil {
    	fmt.Printf(err.Error())
  	} else {
			fmt.Println("File saved!")
		}
	}
}
