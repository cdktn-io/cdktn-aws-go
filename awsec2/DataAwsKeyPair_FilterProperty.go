package awsec2


// Experimental.
type DataAwsKeyPair_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/key_pair#name DataAwsKeyPair#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/key_pair#values DataAwsKeyPair#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

