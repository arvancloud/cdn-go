package main

import (
	"context"
	"log"
	"os"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
	"github.com/gookit/gcli/v3"
)

var (
	ctx = context.Background()
)

func main() {
	api, err := arvancloud.New(
		os.Getenv("ARVANCLOUD_API_KEY"),
		arvancloud.Debug(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	app := gcli.NewApp()
	app.Version = arvancloud.VERSION
	app.Desc = "ArvanCloud CDN"

	configureArgs(app, api)

	app.Run(nil)
}

func configureArgs(app *gcli.App, api *arvancloud.API) {
	domainArg := &gcli.Argument{
		Name:     "domain",
		Desc:     "Your domain, is required",
		Required: true,
	}

	// Get DNS record

	getDNSRecord := &gcli.Command{
		Name:     "get",
		Desc:     "<info>Get</> a single DNS record",
		Aliases:  []string{"g"},
		Examples: `{$fullCmd} <cyan>domain.ir</> <purple>4j6eff36-8e7b-4c12-a7d1-20804e839a67</>`,
		Func: func(cmd *gcli.Command, _ []string) error {
			GetDNSRecord(
				ctx,
				api,
				cmd.Arg("domain").String(),
				cmd.Arg("id").String(),
			)
			return nil
		},
		Config: func(cmd *gcli.Command) {
			cmd.AddArgument(domainArg)
			cmd.AddArg("id", "The id of DNS record, is required", true)
		},
	}

	// List DNS records

	listDNSRecord := &gcli.Command{
		Name:    "list",
		Desc:    "<info>List</> DNS records",
		Aliases: []string{"l", "ls"},
		Examples: `{$fullCmd} <cyan>domain.ir</>
{$fullCmd} <cyan>domain.ir</> <purple>2</>
{$fullCmd} <cyan>domain.ir</> <purple>1</> <green>2</>`,
		Func: func(cmd *gcli.Command, _ []string) error {
			ListDNSRecords(
				ctx,
				api,
				cmd.Arg("domain").String(),
				cmd.Arg("page").Int(),
				cmd.Arg("per-page").Int(),
			)
			return nil
		},
		Config: func(cmd *gcli.Command) {
			cmd.AddArgument(domainArg)
			cmd.AddArg("page", "The page number for pagination", false)
			cmd.AddArg("per-page", "Number of records in each page", false)
		},
	}

	// Delete DNS record

	deleteDNSRecord := &gcli.Command{
		Name:     "delete",
		Desc:     "<info>Delete</> a single DNS record",
		Aliases:  []string{"d", "del"},
		Examples: `{$fullCmd} <cyan>domain.ir</> <purple>4j6eff36-8e7b-4c12-a7d1-20804e839a67</>`,
		Func: func(cmd *gcli.Command, _ []string) error {
			DeleteDNSRecord(
				ctx,
				api,
				cmd.Arg("domain").String(),
				cmd.Arg("id").String(),
			)
			return nil
		},
		Config: func(cmd *gcli.Command) {
			cmd.AddArgument(domainArg)
			cmd.AddArg("id", "The id of DNS record, is required", true)
		},
	}

	app.Add(
		getDNSRecord,
		listDNSRecord,
		deleteDNSRecord,
	)
}
