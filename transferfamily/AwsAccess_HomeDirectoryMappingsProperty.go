package transferfamily


// Experimental.
type AwsAccess_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#entry AwsAccess#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#target AwsAccess#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

