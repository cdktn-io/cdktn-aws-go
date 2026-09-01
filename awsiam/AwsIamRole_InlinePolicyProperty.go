package awsiam


// Experimental.
type AwsIamRole_InlinePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_role#name AwsIamRole#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_role#policy AwsIamRole#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

