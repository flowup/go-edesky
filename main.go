package main

import (
	"github.com/flowup/go-edesky/edesky"
	"fmt"
)

func main() {
	client := edesky.NewClient("")

	dashboards, errs := client.FindDocuments(edesky.DocumentOptions{
		Keywords: "prodej",
	})
	fmt.Println(dashboards)
	fmt.Println(errs)
}
