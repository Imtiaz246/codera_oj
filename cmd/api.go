package cmd

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/handler/http"
	"github.com/spf13/cobra"
	"log"
	goHttp "net/http"
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
	r := http.NewRouter()

	err := goHttp.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}

	return nil
}
