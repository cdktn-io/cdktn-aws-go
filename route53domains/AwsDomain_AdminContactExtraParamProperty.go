package route53domains


// Experimental.
type AwsDomain_AdminContactExtraParamProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name AwsDomain#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#value AwsDomain#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

