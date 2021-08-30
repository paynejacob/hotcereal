package cmd

import (
	"fmt"
	"github.com/paynejacob/hotcereal/pkg/providergen"
	"github.com/spf13/cobra"
	"os"
)

var (
	logLevelString string
)

var rootCmd = &cobra.Command{
	Use: "hotcereal",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(&providergen.Command)
}
