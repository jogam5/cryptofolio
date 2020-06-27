/*
This program manages and organizes a crypto based
portfolio that trades in Bitfinex

To do:

1. Read CSV
2. Write new CSV
3. Upse Tickers instead of Ticket
4. Check for a way to update Google SH in batch

*/
package main

import (
	//"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"time"
	//"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"portfolio/client"
	"portfolio/spreadsheet"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//c, bitfinex := client.ConnectionBitfinex()
	bitfinex := rest.NewClient()

	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	sheet, _ := sh.SheetByTitle("Bitfinex")
	//checkError(err)
	pairsBFX := spreadsheet.FetchPairs(sheet, "BFX")

	/* Bitfinex */
	symbols := make([]string, len(pairsBFX))
	for _, v := range pairsBFX {
		pair := "t" + v.Coin + v.Base
		symbols = append(symbols, pair)
	}

	//data, _ := (bitfinex.Tickers.GetMulti(symbols))
	// 1. Place this information inside a txt
	// We will then find the price information from this txt
	// 2. Create a Map

	// Dereferencing 'data' with *
	//for _, v := range *data {
	//	log.Println(v.LastPrice)
	//}

	//i := 1
	for _, v := range pairsBFX {
		pair := "t" + v.Coin + v.Base
		log.Println(pair)
		r, e := bitfinex.Tickers.Get(pair) // Get this value from TXT, assume it works
		checkError(e)
		sheet.Update(v.Row, 12, spreadsheet.ToS(r.LastPrice))
		//i += 1
		//if i%29 == 0 {
		//	time.Sleep(60 * time.Second)
		//}
	}
	today := time.Now()
	sheet.Update(0, 19, today.Format("01-02-2006 15:04:05"))
	sheet.Synchronize()

	//trades := readCsv()
	//writeCsv(trades, sheet)
	//openSheet(trades)
	//moveTrades(17, sheet, sheetTo)

	//symbols := []string{bitfinex.TradingPrefix + bitfinex.BTCUSD, bitfinex.TradingPrefix + bitfinex.EOSBTC}
	//tickers, err := c.Tickers.GetMulti(symbols)
	//if err != nil {
	//	log.Fatalf("getting ticker: %s", err)
	//}
	//log.Print(tickers)
}
