package arvancloud

type Resource struct {
	Domain string
}

func ResourceDomain(domain string) Resource {
	return Resource{
		Domain: domain,
	}
}
