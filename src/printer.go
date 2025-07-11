package src

import (
	"context"
	"fmt"
	"logs-analyser/pkg/models"
	"os"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Prints the filtered log entries to the console
func Printer(lines *chan models.LineMap, wg *sync.WaitGroup, printableFields []*string, ctx context.Context) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Printer shutting down")
		return
	default:
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(makeHeader(printableFields))

		for l := range *lines {
			t.AppendRow(makeRow(printableFields, &l))
		}
		t.Render()
	}
}

func makeHeader(printableFields []*string) table.Row {
	r := table.Row{}
	for _, f := range printableFields {
		r = append(r, *f)
	}
	return r
}

func makeRow(printableFields []*string, line *models.LineMap) table.Row {
	row := table.Row{}

	for _, fld := range printableFields {
		if val := line.GetField(fld); val != nil {
			row = append(row, *val)
		}
	}

	return row
}
