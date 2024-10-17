package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var c Config
	flag.StringVar(&(c.SourceFilename), "src", "", "source file (*.json)")
	flag.StringVar(&(c.DestFilename), "dst", "", "dest file (*.json)")
	flag.Parse()
	checkError(run(c))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	SourceFilename string
	DestFilename   string
}

func run(c Config) error {
	data, err := ioutil.ReadFile(c.SourceFilename)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = json.Indent(&b, data, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.DestFilename, b.Bytes(), 0655)
}
