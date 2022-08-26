package main

import (
	"time"

	"github.com/oswaldom-code/load-data-stock-donation-project/cmd"
)

func main() {
	initialTime := time.Now()
	cmd.Execute()
	finalTime := time.Now()
	diff := finalTime.Sub(initialTime)
	println(diff.String())
}
