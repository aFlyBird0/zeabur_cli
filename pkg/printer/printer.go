package printer

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type defaultPrinter struct {
}

func New() Printer {
	return &defaultPrinter{}
}

func (p *defaultPrinter) Table(header []string, rows [][]string) {
	columnsCount := len(header)

	colors := []tablewriter.Colors{
		{tablewriter.FgHiMagentaColor},
		{tablewriter.FgGreenColor},
		{tablewriter.FgHiBlueColor},
		{tablewriter.FgHiYellowColor},
		{tablewriter.FgHiCyanColor},
	}
	headerColor := tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetColumnSeparator("")

	headerColors := make([]tablewriter.Colors, columnsCount)
	for i := 0; i < columnsCount; i++ {
		headerColors[i] = headerColor
	}
	table.SetHeaderColor(headerColors...)

	columnColors := make([]tablewriter.Colors, columnsCount)
	for i := 0; i < columnsCount; i++ {
		columnColors[i] = colors[i%len(colors)]
	}

	table.SetColumnColor(columnColors...)

	table.AppendBulk(rows)

	table.Render()
}

var _ Printer = &defaultPrinter{}
