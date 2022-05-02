package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lolchamp",
	Short: "lolchamp is a tool to choose Champion skins set for your friends' party",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
