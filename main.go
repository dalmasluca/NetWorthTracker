package main

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

type NetWorth struct {
	Ammount int
	Date    time.Time
}

func main() {
	nw := NetWorth{Ammount: 20, Date: time.Now()}
	f, err := os.Create("file/NW.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	if err := enc.Encode(nw); err != nil {
		log.Fatal(err)
	}
}
