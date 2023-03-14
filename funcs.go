package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"time"
)

type NetWorth struct {
	Amount float64
	Date   time.Time
}

func WriteNetWorth(nw NetWorth) {
	flog, err := os.OpenFile("log.log", os.O_APPEND, fs.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()
	log.SetOutput(flog)
	f, err := os.OpenFile("file/nw.gob", os.O_APPEND, fs.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	if err := enc.Encode(nw); err != nil {
		log.Fatal(err)
	}
}

func CreateNW() {
	var nw NetWorth
	fmt.Printf("insert your networth: ")
	fmt.Scan(&nw.Amount)
	nw.Date = time.Now()
	WriteNetWorth(nw)
}

func Initilization() {
	os.Create("log.log")
	os.Create("file/nw.gob")
	CreateNW()
	MainLoop()
}

func showMenu() {
	fmt.Println("\t\t\t+----------------------------------+")
	fmt.Println("\t\t\t|                MENU              |")
	fmt.Println("\t\t\t+----------------------------------+")
	fmt.Println("\t\t\t| 1 - show NetWorth                |")
	fmt.Println("\t\t\t+----------------------------------+")
	fmt.Println("\t\t\t| 2 - add to NetWorth              |")
	fmt.Println("\t\t\t+----------------------------------+")
	fmt.Println("\t\t\t| 3 - quit                         |")
	fmt.Println("\t\t\t+----------------------------------+")
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getNW() []NetWorth {
	f, err := os.Open("file/nw.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var nw []NetWorth
	dec := gob.NewDecoder(f)

	if err := dec.Decode(&nw); err != nil {
		if err == io.EOF {
		} else {
			log.Fatal(err)
		}
	}
	return nw
}

func UpdateNetWorth() {
	nw := getNW()
	fmt.Printf("nw: %v\n", nw)
}

func ShowNetWorth() {
	UpdateNetWorth()

}

func MainLoop() {
	var ris = 0
	for ris != 3 {
		cancella()
		showMenu()
		fmt.Scanf("%d", &ris)
		fmt.Print(ris)
		switch ris {
		case 1:
			ShowNetWorth()
		case 2:
			UpdateNetWorth()
		default:
			cancella()
		}
	}
}
