package src

import (
	"bufio"
	"context"
	"os"
	"sync"
)

// Scraper reads lines from a log file.
//
// Parameters:
//
//	path: A pointer to a string containing the path to the log file
//	rawLogLines: A pointer to a channel where each line from the log file will be sent as a string
func Scraper(path *string, rawLogLines *chan string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	defer close(*rawLogLines)

	select {
	case <-ctx.Done():
		return
	default:
		f, err := os.Open(*path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			*rawLogLines <- scanner.Text()
		}
	}
}
