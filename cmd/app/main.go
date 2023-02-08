package main

import (
	"github.com/ProSt1ll/wb-l0/internal/app"
	"log"
)

func main() {
	wb := app.New()
	log.Fatal(wb.Run())
}
