package main

//go:generate gotext -srclang=en update -out=catalogue.go -lang=en,es

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.EuropeanSpanish)
	p.Printf("Hello world!")
	p.Println()
}
