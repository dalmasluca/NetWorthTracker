package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type date struct {
	Year, Month, Day int
}

type NwElements struct {
	Amount float32
	Date   date
}

type Nw []NwElements

func (nw Nw) add(amount float32) Nw {
	date := date{time.Now().Year(), int(time.Now().Month()), time.Now().Day()}
	return append(nw, NwElements{amount, date})
}

func LastDay(year, month int) int {
	switch month {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year%4 == 0 {
			return 29
		}
		return 28
	default:
		return 31
	}
}

func (nw Nw) modify(year, month int, amount float32) error {
	for _, nwElement := range nw {
		if nwElement.Date.Year == year && nwElement.Date.Month == month {
			nwElement.Amount = amount
			break
		}
	}
	err := nw.save("file/nw.json")
	if err != nil {
		log.Printf("in modify function: ")
		log.Fatal(err)
		return err
	}
	return nil
}

func (nw Nw) addSpecific(year, month int, amount float32) Nw {
	for i, element := range nw {
		if element.Date.Year == year && element.Date.Month == month {
			element.Amount = amount
			return nw
		}
		if element.Date.Year == year && element.Date.Month > month {
			newNwElement := NwElements{Amount: amount, Date: date{Year: year, Month: month, Day: LastDay(year, month)}}
			seconPart := Nw{newNwElement}
			seconPart = append(seconPart, nw[i:]...)
			nw = append(nw[:i], seconPart...)
		}
	}
	return nw
}

func (nw Nw) save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	if err := enc.Encode(nw); err != nil {
		return err
	}
	return nil
}

func (nw Nw) load(filename string) (Nw, error) {
	f, err := os.Open(filename)
	if err != nil {
		return *new(Nw), err
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&nw); err != nil {
		return *new(Nw), err
	}
	return nw, nil
}
