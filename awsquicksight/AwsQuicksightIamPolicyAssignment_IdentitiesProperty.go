package awsquicksight


// Experimental.
type AwsQuicksightIamPolicyAssignment_IdentitiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_iam_policy_assignment#group AwsQuicksightIamPolicyAssignment#group}.
	// Experimental.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_iam_policy_assignment#user AwsQuicksightIamPolicyAssignment#user}.
	// Experimental.
	User *[]*string `field:"optional" json:"user" yaml:"user"`
}

