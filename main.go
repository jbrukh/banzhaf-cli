package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"os"

	"github.com/jbrukh/go-banzhaf"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "banzhaf",
		Usage: "calculate Banzhaf power indices",
		Action: func(c *cli.Context) error {
			n := c.NArg()
			if n < 1 {
				return fmt.Errorf("usage: banzhaf-cli <file>")
			}

			args := make([]uint64, 0, 128)

			file, err := os.Open(c.Args().Get(0))
			if err != nil {
				return err
			}
			reader := bufio.NewReader(file)
			scanner := bufio.NewScanner(reader)

			// Set the split function for the scanning operation.
			scanner.Split(bufio.ScanWords)

			// scan
			for scanner.Scan() {
				tok := scanner.Text()
				if v, err := strconv.ParseUint(tok, 10, 64); err != nil {
					return err
				} else {
					args = append(args, v)
				}
			}
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("reading input: %v", err)
			}

			quota := args[0]
			weights := args[1:]

			bi, err := banzhaf.Banzhaf(weights, quota, false)
			if err != nil {
				return err
			}
			for i, v := range bi {
				fmt.Printf("%d,%.20f\n", weights[i], v)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
