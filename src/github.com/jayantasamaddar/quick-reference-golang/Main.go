package main

import "fmt"

const (
    isAdmin = 1 << iota					// 2^0 * 2^0 	= 1
    isHeadquarters						// 2^0 * 2^1 	= 2
    canSeeFinancials					// 2^0 * 2^2 	= 4

    canSeeAfrica						// 2^0 * 2^3 	= 8
    canSeeAsia							// 2^0 * 2^4 	= 16
    canSeeEurope						// 2^0 * 2^5 	= 32
    canSeeNorthAmerica					// 2^0 * 2^6 	= 64
    canSeeSouthAmerica					// 2^0 * 2^7 	= 128
)

func main() {
	/**
	 *  isAdmin 														= 00000001
	 *	canSeeFinancials 												= 00000100
	 *	canSeeEurope													= 00100000
	 *  isAdmin | canSeeFinancials | canSeeEurope 						= 00100101
	*/
    var roles byte = isAdmin | canSeeFinancials | canSeeEurope
    fmt.Printf("%b\n", roles)											// 100101
	fmt.Printf("isAdmin: %v\n", isAdmin & roles == isAdmin)				// isAdmin: true
	fmt.Printf("canSeeAsia: %v\n", canSeeAsia & roles == canSeeAsia)	// canSeeAsia: false
}