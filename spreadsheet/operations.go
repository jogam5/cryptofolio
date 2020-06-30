package spreadsheet

import (
	"bufio"
	"cryptofolio/models"
	"encoding/csv"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

/*
==
Find csv file in 'downloads' folder
==
*/
func FindTradesFile() string {
	path := "/Users/gabriel/downloads/"
	var result string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "jogam6") {
			result = file.Name()
		}
	}
	return result
}

/*
==
Return a struct of Trade objects after
parsing a CSV file
==
*/
func ReadCsv() []models.Trade {
	/* 1. Fetch trades from CSV */
	path := "/Users/gabriel/downloads/"
	trades := []models.Trade{}
	file := FindTradesFile()
	log.Println(file)
	file = filepath.Join(path, file)
	fileReader, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(fileReader))
	id := 1
	/* 2. Loop over trades */
	for {
		line, err := reader.Read()
		log.Println(line)
		if err == io.EOF {
			break
		}
		log.Println(line)
		/* Skip header of file */
		if line[0] != "#" {
			/* 3. Store trades */
			var t models.Trade
			t = models.Trade{
				Id:       id,
				Coin:     line[1][0:3],
				Base:     line[1][4:7],
				Exchange: "BFX",
				Units:    line[2],
				BuyRate:  line[3],
				UUID:     line[0],
				BuyFee:   line[4],
				Date:     line[6],
				Status:   "BUY",
			}
			trades = append(trades, t)
			id = id + 1
		}
	}
	sort.Sort(sort.Reverse(models.ById(trades)))
	return trades
}

/*
==
Receive trades objects and write them in spreadsheet
==
*/
func WriteCsv(trades []models.Trade, sheet *spreadsheet.Sheet) {
	beginRow := int(ReturnLastCell(0, sheet).Row) + 1
	for _, trade := range trades {
		toFind := trade.UUID
		sign := ToF(trade.Units)
		if sign > 0 {
			/* Only BUY trades */
			found := false
			for _, row := range sheet.Rows {
				for _, cell := range row {
					if cell.Value == toFind {
						found = true
						log.Println("Trade already stored")
					}
				}
			}
			if found == false {
				/* Write new trade */
				sheet.Update(beginRow, 0, trade.Coin)
				sheet.Update(beginRow, 1, trade.Base)
				sheet.Update(beginRow, 2, trade.Exchange)
				sheet.Update(beginRow, 3, trade.Date)
				sheet.Update(beginRow, 4, trade.Status)
				sheet.Update(beginRow, 5, trade.Units)
				sheet.Update(beginRow, 6, trade.BuyRate)
				sheet.Update(beginRow, 7, trade.UUID)
				sheet.Update(beginRow, 8, trade.BuyFee)
				beginRow += 1
			}
		}
	}
	today := time.Now()
	sheet.Update(0, 19, today.Format("01-02-2006 15:04:05"))
	sheet.Synchronize()
}
