package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const backoff = "backoff"

// RootCmd configures the command params of sma
var RootCmd = &cobra.Command{
	Use:     "brokerperf",
	Short:   "Simple test harness",
	Version: "1.0.0",
	RunE:    run,
}

func bindOrPanic(key string, flag *pflag.Flag) {
	if err := viper.BindPFlag(key, flag); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().StringP("url", "u", "", "the url to hit")
	RootCmd.Flags().IntP("threads", "t", 1, "number of threads")
	RootCmd.Flags().StringP("apikey", "a", 1, "the api key") // explicit as viper stringslice don't like spaces
	RootCmd.Flags().IntP("iterations", "i", 1, "number of iterations per thread")
	RootCmd.Flags().StringSliceP("header", "H", []string{}, "headers")

	bindOrPanic("url", RootCmd.Flags().Lookup("url"))
	bindOrPanic("apikey", RootCmd.Flags().Lookup("apikey"))
	bindOrPanic("threads", RootCmd.Flags().Lookup("threads"))
	bindOrPanic("iterations", RootCmd.Flags().Lookup("iterations"))
	bindOrPanic("header", RootCmd.Flags().Lookup("header"))
}

// initConfig sets up viper to read from env vars
func initConfig() {
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvPrefix("perf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func configFromViper(cmd *cobra.Command) *Config {
	if viper.GetString("url") == "" {
		cmd.Help()
		fmt.Printf("\nError: url is required.\n\n")
		os.Exit(1)
	}

	return &Config{
		URL:        viper.GetString("url"),
		APIKey:     viper.GetString("apikey"),
		Threads:    viper.GetInt("threads"),
		Iterations: viper.GetInt("iterations"),
		Headers:    viper.GetStringSlice("header"),
	}
}

func run(cmd *cobra.Command, _ []string) error {
	config := configFromViper(cmd)

	fmt.Printf("Test config: %s\n", config)
	perf(config)
	return nil
}
