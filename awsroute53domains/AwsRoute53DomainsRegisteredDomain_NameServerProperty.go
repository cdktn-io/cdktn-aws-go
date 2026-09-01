package awsroute53domains


// Experimental.
type AwsRoute53DomainsRegisteredDomain_NameServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#name AwsRoute53DomainsRegisteredDomain#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#glue_ips AwsRoute53DomainsRegisteredDomain#glue_ips}.
	// Experimental.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
}

