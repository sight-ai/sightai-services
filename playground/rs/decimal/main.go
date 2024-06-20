package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	wei, _ := decimal.NewFromString("1")
	d := int8(18)

	dom := decimal.New(10, int32(-1*d))
	fmt.Println(dom)

	res := wei.Mul(dom)
	fmt.Println(res)
}
