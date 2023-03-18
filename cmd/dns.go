package main

import (
	"context"
	"log"
	"os"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
)

func Get(ctx context.Context, domain string, id string) {
	api, err := arvancloud.New(
		os.Getenv("ARVANCLOUD_API_KEY"),
		arvancloud.Debug(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.GetDNSRecord(ctx, resource, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}
