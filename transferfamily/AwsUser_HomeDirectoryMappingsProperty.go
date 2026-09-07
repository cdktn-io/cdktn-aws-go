package transferfamily


// Experimental.
type AwsUser_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#entry AwsUser#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#target AwsUser#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

