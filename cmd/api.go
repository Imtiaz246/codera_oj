package cmd

import (
	"github.com/imtiaz246/codera_oj/app/router"
	"github.com/spf13/cobra"
	"log"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "codera_oj api server",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		err := runApiServer()
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func runApiServer() error {
	app, err := router.New()
	if err != nil {
		return err
	}

	app.Listen(":3000")

	return nil
}
