package main

import (
	"fmt"

	"github.com/ChinmayR/PackageB"
	"github.com/ChinmayR/PackageC"
)

func main() {
	fmt.Print("From Package A :" + PackageB.FuncInPackageB() +
		PackageC.FuncInPackageC())
}
