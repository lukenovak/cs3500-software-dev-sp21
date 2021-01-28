package main

import (
	"Ormegland/A2/internal/numJson"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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

	stdinStream := json.NewDecoder(os.Stdin)

	_, err = numJson.ParseNumJsonFromStream(stdinStream)
	if err != nil {
		fmt.Println(err.Error())
	}
	// TODO: Full functionality here, calling to internal package
	if *addFlag {
		fmt.Println("now we add")
	} else {
		fmt.Println("now we multiply")
	}

	fmt.Printf("internal has successfully run")
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
