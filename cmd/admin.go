/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin command is used for admin management",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
