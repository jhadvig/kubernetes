package buildconfig

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/buildconfig/buildconfigapi"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubecfg"
)

var buildColumns = []string{"ID", "Type", "SourceURI"}

func RegisterPrintHandlers(printer *kubecfg.HumanReadablePrinter) {
	printer.Handler(buildColumns, printBuildConfig)
	printer.Handler(buildColumns, printBuildConfigList)
}

func printBuildConfig(bc *buildconfigapi.BuildConfig, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\t%s\t%s\n", bc.ID, bc.Type, bc.SourceURI)
	return err
}
func printBuildConfigList(buildList *buildconfigapi.BuildConfigList, w io.Writer) error {
	for _, buildConfig := range buildList.Items {
		if err := printBuildConfig(&buildConfig, w); err != nil {
			return err
		}
	}
	return nil
}
