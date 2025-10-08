package cli

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"

	gotoolsmigrator "github.com/Arthur1/go-tools-migrator"
	"github.com/Arthur1/go-tools-migrator/internal/gotool"
	"github.com/alecthomas/kong"
)

type cli struct {
	Version     bool   `name:"version" short:"v" help:"Print version and quit"`
	DryRun      bool   `name:"dryrun" help:"Output the contents of the new go.mod without modifying existing files."`
	ToolsGoFile string `name:"tools-go-file" default:"tools.go" help:"tools.go file path (default: tools.go)"`
	GoModFile   string `name:"go-mod-file" default:"go.mod" help:"go.mod file path (default: go.mod)"`
}

func (c *cli) Run() error {
	if c.Version {
		return printVersion()
	}

	result, err := gotool.Migrate(c.ToolsGoFile, c.GoModFile, c.DryRun)
	if err != nil {
		return err
	}

	if c.DryRun {
		fmt.Print(result)
	} else {
		fmt.Fprintf(os.Stderr, "✅ Succeeded to migrate.\n")
	}
	return nil
}

type Cli struct{}

func (c *Cli) Run() {
	kctx := kong.Parse(new(cli),
		kong.Name("go-tools-migrator"),
		kong.Description("go-tools-migrator: a CLI tool that replaces tools management via tools.go with go.mod tool directive."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
	)
	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}

func printVersion() error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintf(writer, "go-tools-migrator: a CLI tool that replaces tools management via tools.go with go.mod tool directive.\n") //nolint:errcheck
	fmt.Fprintf(writer, "Version:\t%s\n", gotoolsmigrator.Version)                                                                 //nolint:errcheck
	fmt.Fprintf(writer, "Go version:\t%s\n", runtime.Version())                                                                    //nolint:errcheck
	fmt.Fprintf(writer, "Arch:\t%s\n", runtime.GOARCH)                                                                             //nolint:errcheck
	return writer.Flush()
}
