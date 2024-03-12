/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nambrosini/kcp/config"
	"github.com/nambrosini/kcp/kafka"

	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consumes messages from a topic",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		kw := kafka.NewKafkaReader(cfg.BootstrapServer, cfg.GroupId, topic, cfg.KafkaCertificate, cfg.Username, cfg.Password)
		kw.Read()
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	consumeCmd.Flags().StringVarP(&topic, "topic", "t", "", "sets the topic the consumer reads from")
	consumeCmd.MarkFlagRequired("topic")
}
