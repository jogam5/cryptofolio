package client

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"log"
)

func ConnectionGoogle(spreadsheetName string) spreadsheet.Spreadsheet {
	data, err := ioutil.ReadFile("go-cry-ade01b0b9a7c.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetName)
	//spreadsheet, err := service.FetchSpreadsheet("1yLdidIUEIVJNVnSmMTKkALBj76cF8bI_HSGoR0QmFUg")
	checkError(err)
	//sheet, err := spreadsheet.SheetByTitle("Bitfinex")
	//checkError(err)
	//sheetTo, err := spreadsheet.SheetByTitle("Sold-2018")
	//checkError(err)
	//pairsBFX := spreadsh.FetchPairs(sheet, "BFX")
	/* Bitfinex */
	log.Println("Google connected")
	return spreadsheet
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
