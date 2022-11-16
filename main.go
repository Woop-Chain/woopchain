package main

import (
	_ "embed"

	"github.com/Woop-Chain/woopchain/command/root"
	"github.com/Woop-Chain/woopchain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
