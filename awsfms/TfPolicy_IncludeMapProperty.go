package awsfms


// Experimental.
type TfPolicy_IncludeMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#account TfPolicy#account}.
	// Experimental.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#orgunit TfPolicy#orgunit}.
	// Experimental.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

