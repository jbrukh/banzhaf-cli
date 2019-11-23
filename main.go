package main

import (
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
			if n < 2 {
				return fmt.Errorf("provide <quota> <weight1> <weight2> ...")
			}
			args := make([]uint64, n)
			for i, arg := range c.Args().Slice() {
				var err error
				if args[i], err = strconv.ParseUint(arg, 10, 64); err != nil {
					return err
				}
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
