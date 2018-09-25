package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use: "refresh",

	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("refresh %+v\n", c) //TODO()
	},
}
