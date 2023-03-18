package main

import (
	"context"
	"log"

	arvancloud "github.com/arvancloud/cdn-go/pkg"
)

func GetDNSRecord(ctx context.Context, api *arvancloud.API, domain string, id string) {
	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.GetDNSRecord(ctx, resource, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

func ListDNSRecords(ctx context.Context, api *arvancloud.API, domain string, page int, perPage int) {
	param := arvancloud.ListDNSRecordsParams{
		Type:    "a",
		Page:    page,
		PerPage: perPage,
	}

	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.ListDNSRecords(ctx, resource, param)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}

func DeleteDNSRecord(ctx context.Context, api *arvancloud.API, domain string, id string) {
	resource := arvancloud.Resource{
		Domain: domain,
	}

	u, err := api.DeleteDNSRecord(ctx, resource, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	arvancloud.PrettyPrint(u)
}
