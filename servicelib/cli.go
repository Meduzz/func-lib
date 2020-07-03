package servicelib

import (
	"github.com/Meduzz/func-lib/servicelib/commands"
	"github.com/Meduzz/func-lib/servicelib/service"
)

func Run(service *service.ServiceDefinitionDTO) {
	root := commands.Root(service)
	start := commands.Start(service)
	upload := commands.Upload(service)
	info := commands.Info(service)

	root.AddCommand(start, upload, info)
	root.Execute()
}
