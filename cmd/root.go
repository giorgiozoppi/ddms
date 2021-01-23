package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base CLI command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ddms --config /etc/ddms.yaml",
	Short:   "A peer-to-peer publish/subscribe system",
	Long:    `Data Distribution Message Server`,
	Version: "0.0.1",
}

type DDMSConfig struct {
	ServerAddress     string
	ServerPort        int
	ServerNode        bool
	ServerNodeAddress string
	ProtocolPeriod    int
}
var Config *DDMSConfig

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootFlags := rootCmd.PersistentFlags()
	rootFlags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ddms)")
	rootFlags.StringP("local-peer-address", "l", "127.0.0.0", "local address for listening messages")
	rootFlags.Int16P("peer-port", "x", 15000, "local port for listening messages")
	rootFlags.StringP("remote-server-address", "r", "127.0.0.0", "local address for listening messages")
	rootFlags.Int16P("server-port", "z", 15000, "local port for listening messages")
	rootFlags.BoolP("super-node", "s", false, "set the peer as standalone supernode")
	rootFlags.Int32P("swim-timeout", "t", 15000, "local port for listening messages")
	initConfig()
}
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	Config = new(DDMSConfig)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".ddms")
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf(" config file not found ")
		os.Exit(1)
	}
}
