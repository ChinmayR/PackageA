package main

import (
	"fmt"

	"github.com/ChinmayR/PackageB"
	"github.com/ChinmayR/PackageC"
	"github.com/urfave/cli"
)

func main() {
	fmt.Print("From Package A :" + PackageB.FuncInPackageB() +
		PackageC.FuncInPackageC())

	app := cli.NewApp()
	app.Usage = "Build a dependency on it"
}
