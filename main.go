/*
This program manages and organizes a crypto based
portfolio that trades in Bitfinex

*/
package main

import (
	"cryptofolio/client"
	"cryptofolio/spreadsheet"
	"log"
)

func main() {
	_, bitfinex := client.ConnectionBitfinex()
	sh := client.ConnectionGoogle("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	sheet, _ := sh.SheetByTitle("Bitfinex")
	sheetSold, _ := sh.SheetByTitle("Sold")

	pairsBFX := spreadsheet.FetchPairs(sheet, "BFX")
	log.Println(pairsBFX)

	spreadsheet.UpdatePrice(bitfinex, pairsBFX, sheet)
	trades := spreadsheet.ReadCsv()
	spreadsheet.WriteCsv(trades, sheet, sheetSold)
	//spreadsheet.MoveTrades(4, sheet, sheetSold)
	//r, _ := bitfinex.Tickers.Get("tLINK:USD")
	//log.Println(r)
}
