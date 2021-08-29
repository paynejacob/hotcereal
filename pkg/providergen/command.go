package providergen

import "github.com/spf13/cobra"

var Command = cobra.Command{
	Use: "providergen",
	Run: GenerateProvider,
}
