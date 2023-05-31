package main

import (
	"flag"
	"fmt"
	"github.com/djooberlee/terraform-provider-salt/salt"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"io"
	"os"
)

var version = "was not built correctly" // set via the Makefile

func main() {
	versionFlag := flag.Bool("version", false, "print version information and exit")
	flag.Parse()
	if *versionFlag {
		printVersion(os.Stdout)
		os.Exit(0)
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: salt.Provider,
	})
}

func printVersion(writer io.Writer) {
	fmt.Fprintf(writer, "%s %s\n", os.Args[0], version)
}
