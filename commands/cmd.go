package commands

import (
	"fmt"
	"os"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by drprune main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
}
