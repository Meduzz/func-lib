package commands

import (
	"../service"
	"github.com/spf13/cobra"
)

func Start(service *service.ServiceDefinitionDTO) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the service with one of its provided transports",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			selected := cmd.Flags().Arg(0)

			for _, api := range service.APIs {
				if api.Type() == selected {
					api.Start()
					break
				}
			}
		},
	}
}
