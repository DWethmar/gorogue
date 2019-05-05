package main

import (
	"fmt"

	"github.com/DWethmar/gorogue/console"
)

func main() {

	con, err := console.New(10, 10, "gorogue")

	if err != nil {
		panic(err)
	}

	fmt.Printf(con.Title)

	con.Start(1)
}
