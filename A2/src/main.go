package main

import (
	"flag"
	"fmt"
)

const add = "sum"
const addDesc = "If toggled, starts the program in sum mode"
const product = "product"
const productDesc = "If toggled, starts the program in product mode"

func main() {
	addFlag := flag.Bool(add, false, addDesc)
	productFlag := flag.Bool(product, false, productDesc)

	flag.Parse()

	err := verifyFlags(addFlag, productFlag)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	// TODO: Full functionality here, calling to a2 package
	if *addFlag {
		fmt.Println("now we add")
	} else {
		fmt.Println("now we multiply")
	}

	fmt.Printf("a2 has successfully run")
}

// ensures that flags are present and not duplicated
func verifyFlags(addFlag *bool, productFlag *bool) error {
	if !*addFlag && !*productFlag {
		return fmt.Errorf("no run mode selected")
	} else if *addFlag && *productFlag {
		return fmt.Errorf("multiple run modes selected")
	} else {
		return nil
	}
}
