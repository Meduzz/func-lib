package commands

import (
	"../service"
	"github.com/spf13/cobra"
)

func Root(service *service.ServiceDefinitionDTO) *cobra.Command {
	return &cobra.Command{
		Use:     "service",
		Version: service.Version,
	}
}
