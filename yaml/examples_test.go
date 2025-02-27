package yaml

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
)

func ExampleYAML() {
	configFiles := []altsrc.Sourcer{
		altsrc.StringSourcer(filepath.Join(testdataDir, "config.yaml")),
		altsrc.StringSourcer(filepath.Join(testdataDir, "alt-config.yaml")),
	}

	app := &cli.Command{
		Name: "greet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Sources: YAML("greet.name", configFiles...),
			},
			&cli.IntFlag{
				Name:    "enthusiasm",
				Aliases: []string{"!"},
				Sources: YAML("greet.enthusiasm", configFiles...),
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			punct := ""
			if cmd.Int("enthusiasm") > 9000 {
				punct = "!"
			}

			fmt.Fprintf(os.Stdout, "Hello, %[1]v%[2]v\n", cmd.String("name"), punct)

			return nil
		},
	}

	// Simulating os.Args
	os.Args = []string{"greet"}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stdout, "OH NO: %[1]v\n", err)
	}

	// Output:
	// Hello, Berry!
}
