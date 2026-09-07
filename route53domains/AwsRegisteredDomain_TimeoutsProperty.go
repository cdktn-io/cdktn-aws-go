package route53domains


// Experimental.
type AwsRegisteredDomain_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#create AwsRegisteredDomain#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#update AwsRegisteredDomain#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

