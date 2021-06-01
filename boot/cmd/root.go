package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "boot --config /etc/ddms.yaml",
	Short:   "A peer-to-peer publish/subscribe system",
	Long:    `SuperNode Boot for DDMS`,
	Version: "0.0.1",
}
var cfgFile string

const (
	ipAddress       = "ip-address"
	certificateKey  = "certificate-key"
	certificateFile = "certificate-file"
	configFlag      = "config"
	helpFlag        = "help"
	versionFlag     = "version"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, configFlag, "supernode", "config file (default is $HOME/.supernode.yaml)")
	rootCmd.PersistentFlags().StringP(ipAddress, "i", "127.0.0.1", "server IP Address")
	rootCmd.PersistentFlags().StringP(certificateFile, "x", "", "certificate file (default is $HOME/.supernode_crt.crt)")
	rootCmd.PersistentFlags().StringP(certificateKey, "z", "", "certificate file (default is $HOME/.supernode_key.crt)")
	rootCmd.PersistentFlags().BoolP(helpFlag, "h", false, "print information about boot and its commands")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootOnlyFlags := rootCmd.LocalFlags()
	rootOnlyFlags.Bool(versionFlag, false, "show boot version and exit")
	rootCmd.AddCommand(startCommand)
}

func initConfig() {
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
