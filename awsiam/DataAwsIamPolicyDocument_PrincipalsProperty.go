package awsiam


// Experimental.
type DataAwsIamPolicyDocument_PrincipalsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#identifiers DataAwsIamPolicyDocument#identifiers}.
	// Experimental.
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#type DataAwsIamPolicyDocument#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

