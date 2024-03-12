/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/nambrosini/kcp/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kpc",
	Short: "Kafka producer and consumer",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var (
	cfg   config.Config
	topic string
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
