package main

import (
	"context"
	"os"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
	"github.com/gookit/gcli/v3"
)

var (
	API_KEY = ""
	ctx     = context.Background()
)

func main() {
	API_KEY = os.Getenv("ARVANCLOUD_API_KEY")

	app := gcli.NewApp()
	app.Version = arvancloud.VERSION
	app.Desc = "ArvanCloud CDN"

	configureArgs(app)

	app.Run(nil)
}

func configureArgs(app *gcli.App) {
	domainArg := &gcli.Argument{
		Name:     "domain",
		Desc:     "Your domain",
		Required: true,
	}

	// Get DNS record

	getDNSRecord := &gcli.Command{
		Name:    "get",
		Desc:    "<info>Get</> a single DNS record",
		Aliases: []string{"g"},
		Func: func(cmd *gcli.Command, _ []string) error {
			Get(ctx, cmd.Arg("domain").String(), cmd.Arg("id").String())
			return nil
		},
		Config: func(cmd *gcli.Command) {
			cmd.AddArgument(domainArg)
			cmd.AddArg("id", "The id of DNS record, is required", true)
		},
	}

	app.Add(
		getDNSRecord,
	)
}
