package cmd

import (
	"context"
	"fmt"
	"logs-analyser/cmd/utils"
	"logs-analyser/pkg/models"
	"logs-analyser/src"
	"os"

	"github.com/spf13/cobra"
)

// Default path to the log file
var pathToLogFile string

var rawFields string
var PrintableFields []*string

// Represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "logs-analyser",
	Short: "A command-line tool for analysing log files",
	Run: func(cmd *cobra.Command, args []string) {
		PrintableFields = utils.ParseFields(rawFields, PrintableFields)
		src.RunAnalyser(&pathToLogFile, &models.FilterParams{}, PrintableFields, cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&pathToLogFile, "path", "p", `logs\access.txt`, "Path to the log file")
	rootCmd.PersistentFlags().StringVarP(&rawFields, "fields", "f", "IP,RemoteUser,AuthUser,Time,ReqLine,StateCode,Size,Referer,UserAgent", "Fields to print (comma-separated)")
}

func Execute(ctx context.Context, cancel context.CancelFunc) error {
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
		cancel()
		os.Exit(0)
	})

	// Use ExecuteContext to propagate ctx to rootCmd.Run
	return rootCmd.ExecuteContext(ctx)
}
