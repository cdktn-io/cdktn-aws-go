package route53domains


// Experimental.
type AwsDomain_NameServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#glue_ips AwsDomain#glue_ips}.
	// Experimental.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name AwsDomain#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

