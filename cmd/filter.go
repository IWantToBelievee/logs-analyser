package cmd

import (
	"logs-analyser/cmd/utils"
	"logs-analyser/pkg/models"
	"logs-analyser/src"

	"github.com/spf13/cobra"
)

// FParams is an instance of FilterParams that will be used to store filter criteria
var FParams = models.FilterParams{}

// filterCmd represents the command for filtering log files
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filtering log files",
	Run: func(cmd *cobra.Command, args []string) {
		PrintableFields = utils.ParseFields(rawFields, PrintableFields)
		src.RunAnalyser(&pathToLogFile, &FParams, PrintableFields, cmd.Context())
	},
}

func init() {
	filterCmd.Flags().StringVarP(&FParams.IP, "ip", "i", "", "Filter by IP address")
	filterCmd.Flags().StringVarP(&FParams.RemoteUser, "remote-user", "r", "", "Filter by remote user")
	filterCmd.Flags().StringVarP(&FParams.AuthUser, "auth-user", "a", "", "Filter by authenticated user")
	//	filterCmd.Flags().StringVarP(&FParams.Time, "time", "t", "", "Filter by time")
	filterCmd.Flags().StringVarP(&FParams.ReqLine, "req", "q", "", "Filter by request line")
	filterCmd.Flags().IntVarP(&FParams.StateCode, "state-code", "c", 0, "Filter by state code")
	filterCmd.Flags().IntVarP(&FParams.Size, "size", "z", 0, "Filter by size")
	filterCmd.Flags().StringVarP(&FParams.Referer, "referer", "R", "", "Filter by referer")
	filterCmd.Flags().StringVarP(&FParams.UserAgent, "agent", "A", "", "Filter by user agent")

	// Add filter command to the root command
	rootCmd.AddCommand(filterCmd)
}
