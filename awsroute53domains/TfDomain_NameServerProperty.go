package awsroute53domains


// Experimental.
type TfDomain_NameServerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#glue_ips TfDomain#glue_ips}.
	// Experimental.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name TfDomain#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

