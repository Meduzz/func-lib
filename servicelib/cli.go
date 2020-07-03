package servicelib

import (
	"github.com/Meduzz/func-lib/servicelib/commands"
	"github.com/Meduzz/func-lib/servicelib/service"
)

func Run(svc *service.ServiceDefinitionDTO) {
	root := commands.Root(svc)
	start := commands.Start(svc)
	upload := commands.Upload(svc)
	info := commands.Info(svc)

	root.AddCommand(start, upload, info)
	root.Execute()
}
