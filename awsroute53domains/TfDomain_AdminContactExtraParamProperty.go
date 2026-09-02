package awsroute53domains


// Experimental.
type TfDomain_AdminContactExtraParamProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name TfDomain#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#value TfDomain#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

