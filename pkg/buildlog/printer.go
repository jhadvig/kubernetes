package buildlog

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/buildlog/buildlogapi"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubecfg"
)

var buildLogColumns = []string{"Timestamp", "Message"}

func RegisterPrintHandlers(printer *kubecfg.HumanReadablePrinter) {
	printer.Handler(buildLogColumns, printBuildLog)
}

func printBuildLog(bl *buildlogapi.BuildLog, w io.Writer) error {
	for _, log := range bl.LogItems {
		if _, err := fmt.Fprintf(w, "%s\t%s\n", log.Timestamp, log.Message); err != nil {
			return err
		}
	}
	return nil
}
