package main

import (
	"fmt"
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
	RootCmd.Flags().IntP("iterations", "i", 1, "number of iterations per thread")
	RootCmd.Flags().StringSliceP("header", "H", []string{}, "headers")

	bindOrPanic("url", RootCmd.Flags().Lookup("url"))
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

func configFromViper() *Config {
	return &Config{
		URL:        viper.GetString("url"),
		Threads:    viper.GetInt("threads"),
		Iterations: viper.GetInt("iterations"),
		Headers:    viper.GetStringSlice("header"),
	}
}

func run(_ *cobra.Command, _ []string) error {
	config := configFromViper()

	fmt.Printf("Test config: %s\n", config)
	perf(config)
	return nil
}