// range.go range command code
package rangefilter

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tphakala/birdnet-go/internal/conf"
)

// Command creates the range parent command
func Command(settings *conf.Settings) *cobra.Command {
	rangeCmd := &cobra.Command{
		Use:   "range",
		Short: "Commands related to range operations in BirdNET-Go",
	}

	fmt.Print("here")

	// Add subcommands here
	rangeCmd.AddCommand(PrintCommand(settings))

	return rangeCmd
}
