package commands

import (
	"log"

	"github.com/Meduzz/func-lib/servicelib/service"
	"github.com/spf13/cobra"
)

func Start(svc *service.ServiceDefinitionDTO) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the service with one of its provided transports",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			selected := cmd.Flags().Arg(0)
			var candidate service.API

			for _, api := range svc.APIs {
				if api.ApiName() == selected {
					candidate = api
					break
				}
			}

			if candidate == nil {
				log.Fatalf("Transport named %s was not found", selected)
			}

			err := candidate.Start()

			if err != nil {
				panic(err)
			}
		},
	}
}
