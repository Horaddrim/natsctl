package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// NewTable returns a pointer to a new table with default configuration
func NewTable(headers []string) (table *tablewriter.Table) {
	table = tablewriter.NewWriter(os.Stdout)

	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	return
}
