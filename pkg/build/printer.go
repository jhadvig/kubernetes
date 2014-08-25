package build

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/build/buildapi"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubecfg"
)

var buildColumns = []string{"ID", "Status", "Pod ID"}

func RegisterPrintHandlers(printer *kubecfg.HumanReadablePrinter) {
	printer.Handler(buildColumns, printBuild)
	printer.Handler(buildColumns, printBuildList)
}

func printBuild(build *buildapi.Build, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\t%s\t%s\n", build.ID, build.Status, build.PodID)
	return err
}
func printBuildList(buildList *buildapi.BuildList, w io.Writer) error {
	for _, build := range buildList.Items {
		if err := printBuild(&build, w); err != nil {
			return err
		}
	}
	return nil
}
