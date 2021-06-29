package main

import (
	"github.com/dmacvicar/terraform-provider-salt/salt"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

var version = "was not built correctly" // set via the Makefile

func main() {
	versionFlag := flag.Bool("version", false, "print version information and exit")
	flag.Parse()
	if *versionFlag {
		printVersion(os.Stdout)
		os.Exit(0)
	}

	defer libvirt.CleanupLibvirtConnections()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: libvirt.Provider,
	})
}

func printVersion(writer io.Writer) {
	fmt.Fprintf(writer, "%s %s\n", os.Args[0], version)
}
