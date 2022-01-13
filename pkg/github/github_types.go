package github

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pterm/pterm"
)

// ContainerPackage
type ContainerPackage struct {
	ID         int
	Name       string
	Index      int
	Owner      string
	Visibility string
	CreatedAt  time.Time
}

// PrettyPrintContainerPackage
func (c ContainerPackage) PrettyPrintContainerPackage() {
	pterm.DefaultHeader.
		WithBackgroundStyle(pterm.NewStyle(pterm.BgDarkGray)).
		WithTextStyle(pterm.NewStyle(pterm.FgLightWhite)).
		WithMargin(45).
		Println("Package Information")
	pterm.Println()
	versionTable := table.NewWriter()
	versionTable.SetOutputMirror(os.Stdout)
	versionTable.AppendHeader(table.Row{"Info", "Content"})
	versionTable.AppendRows([]table.Row{
		{"➤ ID", c.ID},
		{"➤ Name", c.Name},
		{"➤ Index", c.Index},
		{"➤ Owner", c.Owner},
		{"➤ Visibility", c.Visibility},
		{"➤ Created At", c.CreatedAt},
	})
	versionTable.SetStyle(table.StyleColoredBright)
	versionTable.Render()
	pterm.Println()
}
