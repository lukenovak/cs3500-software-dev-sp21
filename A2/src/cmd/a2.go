package main

import (
	"Ormegland/A2/internal/numJson"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const sumFlag = "sum"
const sumDesc = "If toggled, starts the program in sum mode"
const product = "product"
const productDesc = "If toggled, starts the program in product mode"

func main() {
	addFlag := flag.Bool(sumFlag, false, sumDesc)
	productFlag := flag.Bool(product, false, productDesc)

	flag.Parse()

	err := verifyFlags(addFlag, productFlag)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	stdinStream := json.NewDecoder(os.Stdin)

	numJsons, err := numJson.ParseNumJsonFromStream(stdinStream)

	// quit if there's a non-NumJson input
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	var output json.RawMessage
	if *addFlag {
		output = numJson.GenerateOutput(numJsons, numJson.Sum)
	} else {
		output = numJson.GenerateOutput(numJsons, numJson.Product)
	}
	if output == nil {
		fmt.Println("Error: no input")
		os.Exit(1)
	}
	_, err = os.Stdout.Write(output)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "\n")

	os.Exit(0)
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
