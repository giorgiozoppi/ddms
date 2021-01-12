package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base CLI command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ddms-server",
	Short:   "Data Distribution Message Server",
	Long:    "\nData Distribution Message Servertool.",
	Version: "<unknown>",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once for the rootCmd.
func Execute() {
	configureCommands()
	configureLogger()
	err := rootCmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		os.Exit(1)
	}
}

const (
	configFlag    = "config"
	superNodeFlag = "supernode"
	versionFlag   = "version"
)

func configureCommands() {

	concurrency := runtime.NumCPU()
	if concurrency < 4 {
		concurrency = 4
	}

	// root flags

	// configure cmd

}
