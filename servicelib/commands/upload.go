package commands

import (
	"fmt"
	"net/http"

	"../service"
	"github.com/Meduzz/helper/http/client"
	"github.com/spf13/cobra"
)

var url string

func Upload(service *service.ServiceDefinitionDTO) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload service definition",
		Long:  "Upload service definition to a central service",
		Run: func(cmd *cobra.Command, args []string) {
			req, err := client.POST(url, service)

			if err != nil {
				panic(err)
			}

			res, err := req.Do(http.DefaultClient)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Uploaded to %s with status %d\n", url, res.Code())
		},
	}

	cmd.Flags().StringVar(&url, "url", "", "http path to upload to")
	cmd.MarkFlagRequired("url")

	return cmd
}
