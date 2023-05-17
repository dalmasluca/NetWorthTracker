package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var Asset_Type = []string{"Stock", "Bond", "Liquidity", "Debit", "Credit", "Crypto"}

type Asset struct {
	Name     string
	Ticket   string
	Amount   float32
	Type     string
	Currency string
}

type portfolioElement struct {
	Assets []Asset
	Date   date
}

type portoflio []portfolioElement

func checkAsset_Type(assset string) bool {
	for _, s := range Asset_Type {
		if s == assset {
			return true
		}
	}
	return false
}

func (p portoflio) CreatePortoflio() portoflio {
	var ris string
	var assets []Asset
	var pElement portfolioElement
	for {
		fmt.Print("aggiunig asset: [y/n]")
		fmt.Scan(&ris)
		if ris != "y" {
			break
		}
		var asset Asset
		var t string
		fmt.Print("insert asset name: ")
		fmt.Scan(&asset.Name)
		fmt.Print("insert asset Ticket: ")
		fmt.Scan(&asset.Ticket)
		fmt.Print("inset asset Amount: ")
		fmt.Scan(&asset.Amount)
		fmt.Print("insert asset Type: ")
		fmt.Scan(&t)
		while(!checkAsset_Type(t)){
			fmt.Pprintln("type doesn't match whit type stored, please retry: ")
			fmt.Scan(&t)	
		}
		asset.Ty
		fmt.Print("insert asset Currency: ")
		fmt.Scan(&asset.Currency)
		assets = append(assets, asset)
	}
	pElement.Assets = assets
	pElement.Date = date{Year: time.Now().Year(), Month: int(time.Now().Month()), Day: time.Now().Day()}
	p = append(p, pElement)
	return p
}

func (p portoflio) save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	if err := enc.Encode(p); err != nil {
		return err
	}
	return nil
}

func (p portoflio) load(filename string) (portoflio, error) {
	f, err := os.Open(filename)
	if err != nil {
		return *new(portoflio), err
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&p); err != nil {
		return *new(portoflio), err
	}
	return p, nil
}
