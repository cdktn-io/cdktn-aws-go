package fms


// Experimental.
type AwsPolicy_ExcludeMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#account AwsPolicy#account}.
	// Experimental.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#orgunit AwsPolicy#orgunit}.
	// Experimental.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

