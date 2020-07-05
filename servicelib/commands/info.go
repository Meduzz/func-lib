package commands

import (
	"fmt"
	"strings"

	"github.com/Meduzz/func-lib/servicelib/service"
	"github.com/spf13/cobra"
)

func Info(service *service.ServiceDefinitionDTO) *cobra.Command {
	apis := make([]string, 0)
	envs := fmt.Sprintf("\nCore service:\n%s", strings.Join(service.Envs, ", "))

	for _, api := range service.APIs {
		apis = append(apis, fmt.Sprintf("%s of type %s", api.ApiName(), api.ApiType()))

		if len(api.Envars()) > 0 {
			envars := strings.Join(api.Envars(), ", ")
			envs = fmt.Sprintf("%s\n\nTransport %s:\n%s", envs, api.ApiName(), envars)
		}
	}

	info := fmt.Sprintf("%s (v%s) expects the following ENVs:\n%s\n\nIt has the following transports:\n%s",
		service.Name,
		service.Version,
		envs,
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
