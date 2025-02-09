package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/keep-starknet-strange/garaga/fp"
	"github.com/keep-starknet-strange/garaga/internal/fptower"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Gnark parser CLI"
	app.Usage = "An example CLI for parsing hint input"
	app.Author = "Bacharif"
	app.Version = "1.0.0"
}

func load_e2_from_args(c *cli.Context, pos int) fptower.E2 {
	var x fptower.E2
	var A0, A1 fp.Element
	n := new(big.Int)
	n, _ = n.SetString(c.Args().Get(pos+0), 10)
	A0.SetBigInt(n)
	n, _ = n.SetString(c.Args().Get(pos+1), 10)
	A1.SetBigInt(n)

	x.A0 = A0
	x.A1 = A1
	return x
}

func load_e6_from_args(c *cli.Context, pos int) fptower.E6 {
	var x fptower.E6
	var x0, x1, x2 fptower.E2
	x0 = load_e2_from_args(c, pos)
	x1 = load_e2_from_args(c, pos+2)
	x2 = load_e2_from_args(c, pos+4)

	x.B0 = x0
	x.B1 = x1
	x.B2 = x2

	return x
}

func load_e12_from_args(c *cli.Context, pos int) fptower.E12 {
	var x fptower.E12
	var x0, x1 fptower.E6
	x0 = load_e6_from_args(c, pos)
	x1 = load_e6_from_args(c, pos+6)

	x.C0 = x0
	x.C1 = x1

	return x
}

func main() {
	info()
	app.Action = func(c *cli.Context) error {

		switch c.Args().Get(0) {
		case "e2":
			var z, x, y fptower.E2
			var A0, A1, A2, A3 fp.Element
			n := new(big.Int)
			n, _ = n.SetString(c.Args().Get(2), 10)
			A0.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(3), 10)
			A1.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(4), 10)
			A2.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(5), 10)
			A3.SetBigInt(n)

			x.A0 = A0
			x.A1 = A1
			y.A0 = A2
			y.A1 = A3

			switch c.Args().Get(1) {
			case "add":
				z.Add(&x, &y)
			case "sub":
				z.Sub(&x, &y)
			case "div":
				z.Div(&x, &y)
			case "mul":
				z.Mul(&x, &y)
			}

			z.A0.FromMont()
			z.A1.FromMont()

			fmt.Println(z)
		case "e6":
			var z, x, y fptower.E6
			var A0, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 fp.Element
			n := new(big.Int)
			n, _ = n.SetString(c.Args().Get(2), 10)
			A0.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(3), 10)
			A1.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(4), 10)
			A2.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(5), 10)
			A3.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(6), 10)
			A4.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(7), 10)
			A5.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(8), 10)
			A6.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(9), 10)
			A7.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(10), 10)
			A8.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(11), 10)
			A9.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(12), 10)
			A10.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(13), 10)
			A11.SetBigInt(n)

			x.B0.A0 = A0
			x.B0.A1 = A1
			x.B1.A0 = A2
			x.B1.A1 = A3
			x.B2.A0 = A4
			x.B2.A1 = A5

			y.B0.A0 = A6
			y.B0.A1 = A7
			y.B1.A0 = A8
			y.B1.A1 = A9
			y.B2.A0 = A10
			y.B2.A1 = A11

			switch c.Args().Get(1) {
			case "add":
				z.Add(&x, &y)
			case "sub":
				z.Sub(&x, &y)
			case "mul":
				z.Mul(&x, &y)
			}
			z.B0.A0.FromMont()
			z.B0.A1.FromMont()
			z.B1.A0.FromMont()
			z.B1.A1.FromMont()
			z.B2.A0.FromMont()
			z.B2.A1.FromMont()

			fmt.Println(z)
		case "e12":
			var z, x, y fptower.E12
			var A0, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 fp.Element
			var B0, B1, B2, B3, B4, B5, B6, B7, B8, B9, B10, B11 fp.Element
			n := new(big.Int)
			n, _ = n.SetString(c.Args().Get(2), 10)
			A0.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(3), 10)
			A1.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(4), 10)
			A2.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(5), 10)
			A3.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(6), 10)
			A4.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(7), 10)
			A5.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(8), 10)
			A6.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(9), 10)
			A7.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(10), 10)
			A8.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(11), 10)
			A9.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(12), 10)
			A10.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(13), 10)
			A11.SetBigInt(n)

			n, _ = n.SetString(c.Args().Get(14), 10)
			B0.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(15), 10)
			B1.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(16), 10)
			B2.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(17), 10)
			B3.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(18), 10)
			B4.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(19), 10)
			B5.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(20), 10)
			B6.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(21), 10)
			B7.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(22), 10)
			B8.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(23), 10)
			B9.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(24), 10)
			B10.SetBigInt(n)
			n, _ = n.SetString(c.Args().Get(25), 10)
			B11.SetBigInt(n)

			x.C0.B0.A0 = A0
			x.C0.B0.A1 = A1
			x.C0.B1.A0 = A2
			x.C0.B1.A1 = A3
			x.C0.B2.A0 = A4
			x.C0.B2.A1 = A5
			x.C1.B0.A0 = A6
			x.C1.B0.A1 = A7
			x.C1.B1.A0 = A8
			x.C1.B1.A1 = A9
			x.C1.B2.A0 = A10
			x.C1.B2.A1 = A11

			y.C0.B0.A0 = B0
			y.C0.B0.A1 = B1
			y.C0.B1.A0 = B2
			y.C0.B1.A1 = B3
			y.C0.B2.A0 = B4
			y.C0.B2.A1 = B5
			y.C1.B0.A0 = B6
			y.C1.B0.A1 = B7
			y.C1.B1.A0 = B8
			y.C1.B1.A1 = B9
			y.C1.B2.A0 = B10
			y.C1.B2.A1 = B11

			switch c.Args().Get(1) {
			case "add":
				z.Add(&x, &y)
			case "sub":
				z.Sub(&x, &y)
			case "mul":
				z.Mul(&x, &y)
			}
			z.C0.B0.A0.FromMont()
			z.C0.B0.A1.FromMont()
			z.C0.B1.A0.FromMont()
			z.C0.B1.A1.FromMont()
			z.C0.B2.A0.FromMont()
			z.C0.B2.A1.FromMont()
			z.C1.B0.A0.FromMont()
			z.C1.B0.A1.FromMont()
			z.C1.B1.A0.FromMont()
			z.C1.B1.A1.FromMont()
			z.C1.B2.A0.FromMont()
			z.C1.B2.A1.FromMont()

			fmt.Println(z)

		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
