package servicelib

import (
	"./commands"
	"./service"
)

func Run(service *service.ServiceDefinitionDTO) {
	root := commands.Root(service)
	start := commands.Start(service)
	upload := commands.Upload(service)
	info := commands.Info(service)

	root.AddCommand(start, upload, info)
	root.Execute()
}
