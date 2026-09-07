package iam


// Experimental.
type AwsRole_InlinePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_role#name AwsRole#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_role#policy AwsRole#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

