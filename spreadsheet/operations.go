package spreadsheet

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"portfolio/models"
	"time"
	//"gabriel/spreadsh"
	//"github.com/bitfinexcom/bitfinex-api-go/v1"
	//"golang.org/x/net/context"
	//"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io"
	"os"
	"path/filepath"
)

func ReadCsv() []models.Trade {
	/* 1. Fetch trades from CSV */
	trades := []models.Trade{}
	path := "/Users/gabo/downloads/"
	today := time.Now().Format("2006-01-02") + "-trades.csv"
	f := filepath.Join(path, today)
	file, _ := os.Open(f)
	reader := csv.NewReader(bufio.NewReader(file))
	/* 2. Loop over trades */
	for {
		line, err := reader.Read()
		//fmt.Println(line)
		if err == io.EOF {
			break
		}
		if line[0] == "#" {
			// refactor
			//fmt.Println("go on")
			continue
		} else {
			//fmt.Println("loop")
			/* Store trades */
			var t models.Trade
			t = models.Trade{
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
		}
	}
	return trades
}

func WriteCsv(trades []models.Trade, sheet *spreadsheet.Sheet) {
	/* Write trades once read from CSV */
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
						fmt.Println("Trade already stored")
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
