package awsroute53domains


// Experimental.
type AwsRoute53DomainsDomain_NameServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#glue_ips AwsRoute53DomainsDomain#glue_ips}.
	// Experimental.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name AwsRoute53DomainsDomain#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

