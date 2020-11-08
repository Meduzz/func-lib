package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Meduzz/func-lib/servicelib/service"
	"github.com/Meduzz/helper/http/client"
	"github.com/spf13/cobra"
)

var url string
var print bool

func Upload(service *service.ServiceDefinitionDTO) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload service definition",
		Long:  "Upload service definition to a central service",
		Run: func(cmd *cobra.Command, args []string) {
			if print {
				bs, err := json.Marshal(service)

				if err != nil {
					fmt.Printf("Marshaling definition to json threw error: %v\n", err)
					os.Exit(1)
				}

				fmt.Println(string(bs))
			} else {
				if url == "" {
					fmt.Printf("No url specified (--url <http://...>")
					os.Exit(1)
				}

				req, err := client.POST(url, service)

				if err != nil {
					panic(err)
				}

				res, err := req.Do(http.DefaultClient)

				if err != nil {
					fmt.Printf("Uploading definitions threw error: %v\n", err)
					os.Exit(1)
				}

				fmt.Printf("Uploaded to %s with status %d\n", url, res.Code())
			}
		},
	}

	cmd.Flags().StringVar(&url, "url", "", "http path to upload to")
	cmd.Flags().BoolVar(&print, "print", false, "print result and exit")

	return cmd
}
