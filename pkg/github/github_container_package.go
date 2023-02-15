package github

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pterm/pterm"
)

type ContainerPackage struct {
	ID         int
	Index      int
	Name       string
	Owner      string
	Visibility string
	CreatedAt  time.Time
}

func (cp ContainerPackage) PrintTable() {
	pterm.DefaultHeader.
		WithBackgroundStyle(pterm.NewStyle(pterm.BgDarkGray)).
		WithTextStyle(pterm.NewStyle(pterm.FgLightWhite)).
		WithMargin(45).
		Println("Package Information")

	pterm.Println()

	rows := []table.Row{
		{"➤ ID", cp.ID},
		{"➤ Name", cp.Name},
		{"➤ Index", cp.Index},
		{"➤ Owner", cp.Owner},
		{"➤ Visibility", cp.Visibility},
		{"➤ Created At", cp.CreatedAt},
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Info", "Content"})
	t.AppendRows(rows)
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	pterm.Println()
}
