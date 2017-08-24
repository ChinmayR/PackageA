package main

import (
	"fmt"

	"github.com/ChinmayR/PackageB"
	"github.com/ChinmayR/PackageC"
)

func main() {
	str, _ := PackageC.FuncInPackageC()
	fmt.Print("From Package A :" + PackageB.FuncInPackageB() +
		str)
}
