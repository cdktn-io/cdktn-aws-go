package kms


// Experimental.
type DataAwsSecret_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secret#name DataAwsSecret#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secret#payload DataAwsSecret#payload}.
	// Experimental.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secret#context DataAwsSecret#context}.
	// Experimental.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secret#grant_tokens DataAwsSecret#grant_tokens}.
	// Experimental.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
}

