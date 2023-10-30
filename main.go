package main

import (
	"bangbang/tools"
	"log"
)

func main() {
	arr, exiti := tools.Check()
	log.Println(arr)
	log.Println(exiti)
}
