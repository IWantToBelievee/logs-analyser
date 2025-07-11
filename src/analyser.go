package src

import (
	"context"
	"logs-analyser/pkg/models"
	"sync"
)

// run orchestrates the log processing pipeline by starting the scraper, parser, filter, and printer goroutines.
// It takes the path to the log file and channels for raw, parsed, and filtered log lines
func RunAnalyser(pathToLogFile *string, filterParams *models.FilterParams, printableFields []*string, ctx context.Context) {
	rawLogLines := make(chan string, 1000)              // Channel for raw log lines scraped from the log file
	parsedLogLines := make(chan models.LineMap, 1000)   // Channel for parsed log lines
	filteredLogLines := make(chan models.LineMap, 1000) // Channel for filtered log lines

	// Create a WaitGroup to wait for the completion of all goroutines
	var wg sync.WaitGroup
	wg.Add(4)

	go Scraper(pathToLogFile, &rawLogLines, &wg, ctx)
	go Parser(&rawLogLines, &parsedLogLines, &wg, ctx)
	go Filter(&parsedLogLines, &filteredLogLines, SelectFilters(*filterParams), &wg, ctx)

	if filterParams.IsDefault() {
		go Printer(&parsedLogLines, &wg, printableFields, ctx)
	} else {
		go Printer(&filteredLogLines, &wg, printableFields, ctx)
	}

	wg.Wait()
}
