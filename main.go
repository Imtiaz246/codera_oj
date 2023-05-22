/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/imtiaz246/codera_oj/cmd"
	"github.com/imtiaz246/codera_oj/initializers"
)

func init() {
	err := initializers.Initialize()
	if err != nil {
		panic(err)
	}
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	cmd.Execute()
}
