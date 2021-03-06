package spreadsheet

import (
	"cryptofolio/models"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"strconv"
)

func FetchPairs(sheet *spreadsheet.Sheet, exchange string) []models.Pair {
	/* Fetch pairs of coins */
	log.Println("Fetching data")
	col := sheet.Columns[0]

	/* Loop over pairs */
	pairs := []models.Pair{}
	for _, v := range col {
		if v.Value == "Coin" {
			continue
		} else if v.Value == "Closed" {
			break
		} else if len(v.Value) > 0 {
			base := sheet.Rows[v.Row][1]
			exc := sheet.Rows[v.Row][2]
			if exc.Value == exchange {
				p := models.Pair{
					Row:      int(v.Row),
					Coin:     v.Value,
					Base:     base.Value,
					Exchange: exc.Value,
				}
				pairs = append(pairs, p)
			}
		}
	}
	return pairs
}

/*
==
Find a specific value in the spreadsheet and return a cell
==
*/

func FindValue(sheet *spreadsheet.Sheet, value string) (bool, spreadsheet.Cell) {
	var cellFound spreadsheet.Cell
	var found bool
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Value == value {
				cellFound = cell
				found = true
			}
		}
	}
	return found, cellFound
}

func ReturnLastCell(colNumber uint, sheet *spreadsheet.Sheet) spreadsheet.Cell {
	/* Returns a last not null cell of a specific column */
	var last spreadsheet.Cell
	for _, cell := range sheet.Columns[colNumber] {
		if cell.Value == "" {
			break
		} else {
			last = cell
		}
	}
	return last
}

func ToI(s string) int {
	/* Converts string to int */
	i, _ := strconv.Atoi(s)
	return i
}

func ToF(s string) float64 {
	/* Converts string to float */
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

func ToS(num float64) string {
	/* Convert float64 to string */
	return strconv.FormatFloat(num, 'f', -1, 64)
}
