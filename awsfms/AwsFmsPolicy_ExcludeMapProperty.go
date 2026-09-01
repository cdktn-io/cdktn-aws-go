package awsfms


// Experimental.
type AwsFmsPolicy_ExcludeMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#account AwsFmsPolicy#account}.
	// Experimental.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#orgunit AwsFmsPolicy#orgunit}.
	// Experimental.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

