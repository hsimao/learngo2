package main

import "fmt"

func main() {

	const (
		ETH = 9 - iota
		WAN
		HELLO
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}
	fmt.Println("ETH", ETH)
	fmt.Println("WAN", WAN)
	fmt.Println("HELLO", HELLO)

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("%#v\n", rates)

	gad := [3]string{"hello"}
	fmt.Printf("%q\n", gad)

}
