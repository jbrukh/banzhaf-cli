package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file, f",
				Usage: "specify an input csv file",
			},
		},
		Action: func(c *cli.Context) error {
			file := c.String("file")
			if file != "" {
				return runBanzhaf(file)
			} else {
				n := c.NArg()
				if n < 2 {
					return fmt.Errorf("provide [quota] [weight1 weight2 ...]")
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

				bi, ok := banzhaf.Banzhaf(weights, quota, true)
				if !ok {
					return fmt.Errorf("banzhaf error")
				}
				log.Printf("index=%v\n", bi)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func runBanzhaf(filename string) error {

	log.Printf("%s\n", filename)
	csvFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var (
		weights []uint64
		total   uint64
		quota   uint64
		count   int
	)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		balance, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return err
		}
		balanceUint := uint64(balance)

		if balanceUint > 0 {
			log.Printf("%v\n", balanceUint)
			weights = append(weights, balanceUint)
			total += balanceUint
			count++
		}
	}
	quota = total/2 + 1
	log.Printf("total=%d, quota=%d, count=%d\n", total, quota, count)
	log.Printf("weights=%v\n", weights)

	// do analysis
	bi, ok := banzhaf.Banzhaf(weights, quota, true)
	if !ok {
		return fmt.Errorf("banzhaf error")
	}
	log.Printf("index=%v\n", bi)
	return nil
}
