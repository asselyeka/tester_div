package main

import (
	"fmt"

	student ".."
)

func main() {
	student.PrintNbrBase(919617, "01")
	fmt.Print("\n11100000100001000001 - Passed\n")

	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	fmt.Print("\nHIDAHI - Passed\n")

	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	fmt.Print("\n-MssiD - Passed\n")

	student.PrintNbrBase(-661165, "1")
	fmt.Print("\nNV - Passed\n")

	student.PrintNbrBase(-861737, "Zone01Zone01")
	fmt.Print("\nNV - Passed\n")

	student.PrintNbrBase(125, "0123456789ABCDEF")
	fmt.Print("\n7D - Passed\n")

	student.PrintNbrBase(-125, "choumi")
	fmt.Print("\n-uoi - Passed\n")

	student.PrintNbrBase(125, "-ab")
	fmt.Print("\nNV - Passed\n")

	student.PrintNbrBase(-9223372036854775808, "0123456789")
	fmt.Print("\n-9223372036854775808 - Passed\n")
}

/*
PrintNbrBase
Try with the argument: nbr = 919617 and base = 01

11100000100001000001
Does the function prints the value?
Try with the argument: nbr = 753639 and base = CHOUMIisDAcat!

HIDAHI
Does the function prints the value above?
Try with the argument: nbr = -174336 and base = CHOUMIisDAcat!

-MssiD
Does the function prints the value above?
Try with the argument: nbr = -661165 and base = 1

NV
Does the function prints the value above?
Try with the argument: nbr = -861737 and base = Zone01Zone01

NV
Does the function prints the value above?
Try with the argument: nbr = 125 and base = 0123456789ABCDEF

7D
Does the function prints the value above?
Try with the argument: nbr = -125 and base = choumi

-uoi
Does the function prints the value above?
Try with the argument: nbr = 125 and base = -ab

NV
Does the function prints the value above?
Try with the argument: nbr = -9223372036854775808 and base = 0123456789

-9223372036854775808
Does the function prints the value above?
*/
