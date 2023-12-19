package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

const (
	CLI_NAME    = "op-secrets"
	CLI_VERSION = "v0.0.1"
)

func main() {
	// Global logger flags
	log.SetFlags(0)
	log.SetPrefix(CLI_NAME + ": ")

	app := &cli.App{
		Name:            CLI_NAME,
		Version:         CLI_VERSION,
		HideHelpCommand: true,
		CustomAppHelpTemplate: `
Usage: {{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .Commands}} command{{end}}
{{if .Commands}}
Commands:
{{range .Commands}}  {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}
Options:
{{range .VisibleFlags}}  {{.}}
{{end}}`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "Path to YAML secrets config `file` (required)",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:  "chdir",
				Usage: "Switch to a different `directory` before executing the command",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "Lists secrets in a human-friendly format (redacted)",
				Action: RunCommand,
			},
			{
				Name:   "env",
				Usage:  "Get secrets in .env format",
				Action: RunCommand,
			},
			{
				Name:   "inline",
				Usage:  "Get secrets in an inline format", // Ask Prashant?
				Action: RunCommand,
			},
			{
				Name:   "sh",
				Usage:  "Get secrets in shell export format",
				Action: RunCommand,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func RunCommand(ctx *cli.Context) error {
	if ctx.String("config") == "" {
		return fmt.Errorf("required \"config\" flag not set")
	}
	if ctx.String("chdir") != "" {
		if err := os.Chdir(ctx.String("chdir")); err != nil {
			return err
		}
	}

	secrets, err := ParseYamlConfig(ctx.String("config"))
	if err != nil {
		return err
	}

	// Output result
	switch ctx.Command.Name {
	case "list":
		// Use colors for pretty output
		c1 := color.Yellow.Render
		c2 := color.Cyan.Render
		for _, s := range secrets {
			// Prepare provider string output
			providerStr := c1(s.Provider)
			if s.Vault != "" {
				providerStr = fmt.Sprintf("%s(%s)", c1(s.Provider), c1(s.Vault))
			}
			l := int(len(s.Value) / 3)
			value := fmt.Sprintf("%s%s", s.Value[:l], strings.Repeat("*", len(s.Value)-l))
			fmt.Printf("[%s] %s %s\n", providerStr, c2(s.Key), value)
		}
	case "env":
		for _, s := range secrets {
			fmt.Printf("%s=%s\n", s.Key, s.Value)
		}
	case "inline":
		for _, s := range secrets {
			// Escape special characters
			value := strings.ReplaceAll(s.Value, "'", "'\"'\"'")
			fmt.Printf("%s='%s'\t", s.Key, value)
		}
	case "sh":
		for _, s := range secrets {
			// Escape special characters
			value := strings.ReplaceAll(s.Value, "'", "'\"'\"'")
			fmt.Printf("export %s='%s'\n", s.Key, value)
		}
	}
	return nil
}
