package main

import (
	"encoding/xml"
	"flag"
	"github.com/jackmanlabs/ally-api/lib"
	"github.com/jackmanlabs/errors"
	"log"
	"os"
)

func main() {

	var inputPath *string = flag.String("input", "", "The path of the input file. Required.")
	flag.Parse()

	if *inputPath == "" {
		flag.Usage()
		log.Fatal("A input file path must be specified.")
	}

	inputFile, err := os.Open(*inputPath)
	if os.IsNotExist(err) {
		log.Print("Input file does not exist: ", *inputPath)
		return
	} else if err != nil {
		log.Fatal(errors.Stack(err))
	}
	defer inputFile.Close()

	var accounts lib.AccountsBalancesResponse
	dec := xml.NewDecoder(inputFile)
	err = dec.Decode(&accounts)
	if err != nil {
		log.Print(err)
		log.Print("Failure to read input file.")
		return
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	err = enc.EncodeElement(accounts, xml.StartElement{Name: xml.Name{Local: "response"}})
	if err != nil {
		log.Fatal(errors.Stack(err))
	}

}
