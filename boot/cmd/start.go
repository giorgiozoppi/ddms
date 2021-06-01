package cmd

import (
	"fmt"
	"github.com/giorgiozoppi/ddms/boot/bootserver"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	startUse      = "start"
	startShort    = "Start the boot node"
	startLong     = `Start a boot supernode to a given port that it's been provided by the '`
	startExamples = ` start a supernode: boot start`
)

var startCommand = &cobra.Command{
	Use:     startUse,
	Short:   startShort,
	Long:    startLong,
	Example: startExamples,
	PreRun:  loadConfigs,
	Run:     startCommandMain,
	Args:    cobra.ExactArgs(0),
}

func loadConfigs(cmd *cobra.Command, args []string) {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".supernode")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func startCommandMain(cmd *cobra.Command, args []string) {
	ipAddress := viper.GetString(ipAddress)
	certificateKey := viper.GetString(certificateKey)
	certificateFile := viper.GetString(certificateFile)
	node := bootserver.NewSuperNode(ipAddress, certificateFile, certificateKey)
	node.Run()
}
