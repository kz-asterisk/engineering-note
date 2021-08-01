package main

import (
	"fmt"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	//IPAddr{1, 2, 3, 4} は、 "1.2.3.4"
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println(len(hosts))
	for name, ip := range hosts {
		fmt.Printf("%T: %T\n", name, ip)
		fmt.Printf("%v: %v\n", name, ip)
	}
}
