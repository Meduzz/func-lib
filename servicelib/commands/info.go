package commands

import (
	"fmt"
	"strings"

	"../service"
	"github.com/spf13/cobra"
)

func Info(service *service.ServiceDefinitionDTO) *cobra.Command {
	apis := make([]string, 0)

	for _, api := range service.APIs {
		apis = append(apis, fmt.Sprintf("%s of type %s", api.Name(), api.Type()))
	}

	info := fmt.Sprintf("%s (v%s) expects the following ENVs to be set:\n%s\n\nIt has the following transports:\n%s",
		service.Name,
		service.Version,
		strings.Join(service.Envs, ", "),
		strings.Join(apis, "\n"))

	return &cobra.Command{
		Use:   "info",
		Short: "Display info about the service and its transports",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(info)
		},
	}
}
