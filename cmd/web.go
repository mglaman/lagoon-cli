package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"w"},
	Short:   "Launch the web user interface",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(lagoonCLIConfig.Lagoons[cmdLagoon].UI)
		if lagoonCLIConfig.Lagoons[cmdLagoon].UI != "" {
			urlBuilder.WriteString(fmt.Sprintf("/projects/%s", cmdProjectName))
		} else {
			output.RenderError("unable to determine url for ui, is one set?", outputOptions)
			os.Exit(1)
		}

		url := urlBuilder.String()
		fmt.Println(fmt.Sprintf("Opening %s", url))
		_ = browser.OpenURL(url)
	},
}

var kibanaCmd = &cobra.Command{
	Use:     "kibana",
	Aliases: []string{"k"},
	Short:   "Launch the kibana interface",
	Run: func(cmd *cobra.Command, args []string) {
		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(lagoonCLIConfig.Lagoons[cmdLagoon].Kibana)
		if lagoonCLIConfig.Lagoons[cmdLagoon].Kibana == "" {
			output.RenderError("unable to determine url for kibana, is one set?", outputOptions)
			os.Exit(1)
		}

		url := urlBuilder.String()
		fmt.Println(fmt.Sprintf("Opening %s", url))
		_ = browser.OpenURL(url)
	},
}
