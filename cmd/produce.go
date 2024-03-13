/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nambrosini/kcp/config"
	"github.com/nambrosini/kcp/kafka"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produces a message to a kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		kw := kafka.NewKafkaWriter(cfg.BootstrapServer, topic, cfg.KafkaCertificate, cfg.Username, cfg.Password)
		kw.Produce(key, value)
	},
}

var (
	key   string
	value string
)

func init() {
	rootCmd.AddCommand(produceCmd)

	produceCmd.Flags().StringVarP(&key, "key", "k", "", "the key of the message")
	produceCmd.Flags().StringVarP(&value, "value", "v", "", "the value of the message")
	produceCmd.Flags().StringVarP(&topic, "topic", "t", "", "the topic to write to")
	produceCmd.MarkFlagRequired("topic")
}
